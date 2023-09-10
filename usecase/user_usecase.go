package usecase

import (
	"errors"
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"
	"ewalletgolang/repository"
)

type UserUsecase interface {
	Register(user dto.UserRequest) (entity.User, error)
	IsEmailTaken(email string) (bool, error)
	Login(user dto.UserLoginRequest) (entity.User, error)
	ResetPassword(email, newPassword string) error
	FindUserById(userId int) (entity.User, error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) UserUsecase {
	return &usecase{repository}
}

func (u *usecase) Register(user dto.UserRequest) (entity.User, error) {
	newUser, err := u.repository.Register(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (u *usecase) IsEmailTaken(email string) (bool, error) {
	return u.repository.IsEmailTaken(email)
}

func (u *usecase) Login(user dto.UserLoginRequest) (entity.User, error) {
	userLogin, err := u.repository.FindByEmail(user.Email)
	if err != nil {
		return entity.User{}, errors.New("invalid username or Password")
	}

	errPassword := helper.CheckPasswordHash(userLogin.Password, user.Password)
	if errPassword != nil {
		return entity.User{}, errors.New("invalid username or Password")
	}
	
	return userLogin, nil
}

func (u *usecase) ResetPassword(email, newPassword string) error {
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return err
	}

	return u.repository.UpdatePassword(user, newPassword)
}

func (u *usecase) FindUserById(userId int) (entity.User, error) {
	user, err := u.repository.FindUserById(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}