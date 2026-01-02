package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)

	CreateBankAccount(e domain.BankAccount) error

	// cart
	FindCartItems(uId uint) ([]domain.Cart, error)
	FindCartItem(uId uint, pId uint) (domain.Cart, error)
	CreateCart(c domain.Cart) error
	UpdateCart(c domain.Cart) error
	DeleteCartById(uId uint) error
	DeleteCartItems(uId uint) error

	// order
	CreateOrder(o domain.Order) error
	FindOrders(uId uint) ([]domain.Order, error)
	FindOrderById(id uint, uId uint) (domain.Order, error)

	// Profile
	CreateProfile(e domain.Address) error
	UpdateProfile(e domain.Address) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateOrder implements UserRepository.
func (r *userRepository) CreateOrder(o domain.Order) error {
	err := r.db.Create(&o).Error
	if err != nil {
		return errors.New("order creation failed")
	}
	return nil
}

// FindOrderById implements UserRepository.
func (r *userRepository) FindOrderById(id uint, uId uint) (domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Item").Where("id=? AND user_id=?", id, uId).First(&order).Error
	if err != nil {
		return domain.Order{}, errors.New("unable to find order")
	}
	return order, nil
}

// FindOrders implements UserRepository.
func (r *userRepository) FindOrders(uId uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Where("user_id=?", uId).Find(&orders).Error
	if err != nil {
		return nil, errors.New("unable to find orders")
	}
	return orders, nil
}

// CreateProfile implements UserRepository.
func (r *userRepository) CreateProfile(e domain.Address) error {
	err := r.db.Create(&e).Error
	if err != nil {
		return errors.New("Failed to create profile")
	}
	return nil
}

// UpdateProfile implements UserRepository.
func (r *userRepository) UpdateProfile(e domain.Address) error {
	err := r.db.Where("user_id=?", e.UserId).Updates(&e).Error
	if err != nil {
		return errors.New("Unable to update profile")
	}
	return nil

}

// CreateCart implements UserRepository.
func (r *userRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

// DeleteCartById implements UserRepository.
func (r *userRepository) DeleteCartById(id uint) error {
	return r.db.Delete(&domain.Cart{}, id).Error
}

// DeleteCartItems implements UserRepository.
func (r *userRepository) DeleteCartItems(uId uint) error {
	return r.db.Where("user_id=?", uId).Delete(&domain.Cart{}).Error
}

// FindCartItem implements UserRepository.
func (r *userRepository) FindCartItem(uId uint, pId uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id=? AND product_id=?", uId, pId).Find(&cartItem).Error
	return cartItem, err
}

// FindCartItems implements UserRepository.
func (r *userRepository) FindCartItems(uId uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id=?", uId).Find(&carts).Error
	return carts, err
}

// UpdateCart implements UserRepository.
func (r *userRepository) UpdateCart(c domain.Cart) error {
	var cart domain.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.ID).Updates(c).Error
	return err
}

// CreateBankAccount implements UserRepository.
func (r *userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := r.db.Create(&usr).Error

	if err != nil {
		log.Printf("create user error %v", err)
		return domain.User{}, errors.New("Failed to create user ")
	}

	return usr, nil
}
func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.Preload("Address").First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("User doesn't exist ")
	}
	return user, nil
}
func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	err := r.db.Preload("Address").
		Preload("Cart").
		Preload("Orders").
		First(&user, id).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("User doesn't exist ")
	}
	return user, nil
}
func (r userRepository) UpdateUser(id uint, usr domain.User) (domain.User, error) {

	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(usr).Error
	if err != nil {
		log.Printf("Update user error %v", err)
		return domain.User{}, errors.New("Unable to update user ")
	}
	return user, nil
}
