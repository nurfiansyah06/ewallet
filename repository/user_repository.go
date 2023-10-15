package repository

import (
	"errors"
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Register(user dto.UserRequest) (entity.User, error)
	IsEmailTaken(email string) (bool, error)
	FindByEmail(email string) (entity.User, error)
	ResetPassword(email string, newPassword string) error
	UpdatePassword(user entity.User, newPassword string) error
	FindUserById(userId int) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user dto.UserRequest) (entity.User, error) {
	password, _ := helper.HashPassword(user.Password)
	
	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Password: password, 
	}

	err := r.db.Create(&newUser).Error
	if err != nil {
		log.Fatal(err)
	}

	currentAutoIncrement := 1
	
	newWallet := entity.Wallet{
		WalletNumber: generateWalletNumber(777, currentAutoIncrement+1),
		Amount: 0,
		UserId:       newUser.UserId,
	}

	err = r.db.Create(&newWallet).Error
	if err != nil {
		log.Fatal(err)
	}

	return newUser, err
}

func (r *repository) IsEmailTaken(email string) (bool, error) {
	var user entity.User

	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}

	return true, nil
}

func generateWalletNumber(prefix int, currentAutoIncrement int) int {
	return prefix*1000000000000 + currentAutoIncrement
}

func (r *repository) ResetPassword(email, newPassword string) error {
	user, err := r.FindByEmail(email)
	if err != nil {
		return err
	}

	return r.UpdatePassword(user, newPassword)
}

func (r *repository) UpdatePassword(user entity.User, newPassword string) error {
	return r.db.Model(user).Update("password", newPassword).Error
}

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindUserById(userId int) (entity.User, error) {
	var user entity.User

	err := r.db.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return user, errors.New("no user found on with that ID")
	}

	return user, nil
}