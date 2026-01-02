package service

import (
	"errors"
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/notification"
	"log"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	CRepo  repository.CatelogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s UserService) Signup(input dto.UserSignup) (string, error) {

	hPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}

	log.Println(input)
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hPassword,
		Phone:    input.Phone,
	})

	// generate token

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	// Perform some db opr and business logic

	user, err := s.Repo.FindUser(email)
	return &user, err
}

func (s UserService) Login(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)

	if err != nil {
		return "", errors.New("User doesnt exist")
	}

	err = s.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	// compare password and generate token

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) isVerifiedUser(id uint) bool {

	currentUsr, err := s.Repo.FindUserById(id)

	return err == nil && currentUsr.Verified

}

func (s UserService) GetVerificationCode(e domain.User) error {

	// if user already verified

	if s.isVerifiedUser(e.ID) {
		return errors.New("User already verified")
	}

	// generate verification code

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return nil
	}

	//update user

	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)
	if err != nil {
		return errors.New("Unable to update user")
	}

	user, _ = s.Repo.FindUserById(e.ID)
	// send SMS

	msg := fmt.Sprintf("Your verification code is %v", code)
	notificationClient := notification.NewNotificationClient(s.Config)
	err = notificationClient.SendSms(user.Phone, msg)
	if err != nil {
		return errors.New("Erron on sending notification")
	}
	return nil
}

func (s UserService) VerifyCode(id uint, code int) error {

	if s.isVerifiedUser(id) {
		return errors.New("User already verified")
	}

	user, err := s.Repo.FindUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("Verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("Verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)
	if err != nil {
		return errors.New("Unable to verify user")
	}

	return nil
}

func (s UserService) CreateProfile(id uint, input dto.ProfileInput) error {

	// update user
	var user domain.User

	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}
	_, err := s.Repo.UpdateUser(id, user)

	if err != nil {
		return errors.New("Unable to update user")
	}

	// create address

	address := domain.Address{
		AddressLine1: input.AddressInput.AddressLine1,
		AddressLine2: input.AddressInput.AddressLine2,
		City:         input.AddressInput.City,
		Country:      input.AddressInput.Country,
		PostCode:     input.AddressInput.PostCode,
		UserId:       int(id),
	}
	err = s.Repo.CreateProfile(address)

	if err != nil {
		return errors.New("Unable to create profile")
	}
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	user, err := s.Repo.FindUserById(id)
	if err != nil {
		return nil, errors.New("Unable to get profile")
	}
	return &user, nil
}

func (s UserService) UpdateProfile(id uint, input dto.ProfileInput) error {

	user, err := s.Repo.FindUserById(id)

	if err != nil {
		return err
	}

	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}
	_, err = s.Repo.UpdateUser(id, user)

	if err != nil {
		return errors.New("Unable to update user")
	}
	address := domain.Address{
		AddressLine1: input.AddressInput.AddressLine1,
		AddressLine2: input.AddressInput.AddressLine2,
		City:         input.AddressInput.City,
		Country:      input.AddressInput.Country,
		PostCode:     input.AddressInput.PostCode,
		UserId:       int(id),
	}
	err = s.Repo.UpdateProfile(address)

	if err != nil {
		return errors.New("Unable to update profile")
	}
	return nil
}

func (s UserService) BecomeSeller(id uint, input dto.SellerInput) (string, error) {

	// find exisitng user
	user, _ := s.Repo.FindUserById(id)

	// return already joined seller program

	if user.UserType == domain.SELLER {
		return "", errors.New("You have alreaed joined seller")

	}

	// update user

	seller, err := s.Repo.UpdateUser(id, domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.PhoneNumber,
		UserType:  domain.SELLER,
	})

	if err != nil {
		return "", err
	}

	//gwenerating token
	token, err := s.Auth.GenerateToken(user.ID, user.Email, seller.UserType)

	// create bank account info
	account := domain.BankAccount{
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
		UserId:      id,
	}

	err = s.Repo.CreateBankAccount(account)

	return token, err
}

func (s UserService) FindCart(id uint) ([]domain.Cart, error) {

	carts, err := s.Repo.FindCartItems(id)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s UserService) CreateCart(input dto.CreateCartRequest, u domain.User) ([]domain.Cart, error) {

	// check if cart is exist

	cart, _ := s.Repo.FindCartItem(u.ID, input.ProductId)

	if cart.ID > 0 {
		if input.ProductId == 0 {
			return nil, errors.New("please provide valie product id")
		}
		// delete the cart items
		if input.Qty < 1 {
			err := s.Repo.DeleteCartById(cart.ID)
			if err != nil {
				return nil, errors.New("error on delete cart items")
			}
		} else {
			// update the cart items

			cart.Qty = input.Qty
			err := s.Repo.UpdateCart(cart)
			if err != nil {
				return nil, errors.New("error on update cart item")
			}

		}
	} else {

		// check if product exist
		product, err := s.CRepo.FindProductById(int(input.ProductId))

		if product.ID < 1 {
			return nil, errors.New("Product not found")
		}
		// create cart

		err = s.Repo.CreateCart(domain.Cart{
			ProductId: input.ProductId,
			UserId:    int(u.ID),
			Name:      product.Name,
			ImageUrl:  product.ImageUrl,
			Qty:       input.Qty,
			Price:     product.Price,
			SellerId:  uint(product.UserId),
		})
		if err != nil {
			return nil, errors.New("Error creating cart item")
		}

	}
	return s.Repo.FindCartItems(u.ID)

}

func (s UserService) CreateOrder(u domain.User) (int, error) {

	// find card items
	cartItems, err := s.Repo.FindCartItems(u.ID)

	if err != nil {
		return 0, errors.New("error on finding cart item")
	}

	if len(cartItems) == 0 {
		return 0, errors.New("cart is empty cannot create order")

	}
	//  find success payment
	paymentId := "PAY12345"
	txnId := "TXN12345"
	orderRef, _ := helper.RandomNumbers(8)

	// create order with created order ref
	var amount float64
	var orderItems []domain.OrderItem

	for _, item := range cartItems {
		amount += item.Price * float64(item.Qty)
		orderItems = append(orderItems, domain.OrderItem{
			ProductId: item.ProductId,
			Qty:       item.Qty,
			Price:     item.Price,
			Name:      item.Name,
			ImageUrl:  item.ImageUrl,
			SellerId:  item.SellerId,
		})
	}

	order := domain.Order{
		UserId:         u.ID,
		PaymentId:      paymentId,
		TransactionId:  txnId,
		OrderRefNumber: uint(orderRef),
		Amount:         amount,
		Item:           orderItems,
	}

	err = s.Repo.CreateOrder(order)
	if err != nil {
		return 0, err
	}

	// send email with order details

	// Remove cart items

	err = s.Repo.DeleteCartItems(u.ID)
	log.Printf("error on delete cart items %v", err)

	// return order number
	return orderRef, nil
}

func (s UserService) GetOrders(u domain.User) ([]domain.Order, error) {
	orders, err := s.Repo.FindOrders(u.ID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s UserService) GetOrderById(id uint, uId uint) (domain.Order, error) {
	order, err := s.Repo.FindOrderById(id, uId)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
