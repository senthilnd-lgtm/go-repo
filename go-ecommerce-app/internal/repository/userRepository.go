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
}

type userRepository struct {
	db *gorm.DB
}

// CreateBankAccount implements UserRepository.
func (r *userRepository) CreateBankAccount(e domain.BankAccount) error {
	return r.db.Create(&e).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
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

	err := r.db.First(&user, "email=?", email).Error
	if err != nil {
		log.Printf("Find user error %v", err)
		return domain.User{}, errors.New("User doesn't exist ")
	}
	return user, nil
}
func (r userRepository) FindUserById(id uint) (domain.User, error) {

	var user domain.User

	err := r.db.First(&user, id).Error
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
