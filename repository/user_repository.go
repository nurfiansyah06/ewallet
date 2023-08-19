package repository

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"

	"gorm.io/gorm"
)

type Repository interface {
	Register(user dto.UserRequest) (entity.User, error)
	IsEmailTaken(email string) (bool, error)
	FindByEmail(email string) (entity.User, error)
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
		Email: user.Email,
		Name: user.Name,
		Password: password, 
	}

	err := r.db.Create(&newUser).Error
	
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

func (r *repository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}