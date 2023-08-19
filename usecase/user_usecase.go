package usecase

import (
	"errors"
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"
	"ewalletgolang/middleware"
	"ewalletgolang/repository"
	"os"

	"github.com/joho/godotenv"
)

type Usecase interface {
	Register(user dto.UserRequest) (entity.User, error)
	IsEmailTaken(email string) (bool, error)
	Login(user dto.UserLoginRequest) (string, error)
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) Usecase {
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

func (u *usecase) Login(user dto.UserLoginRequest) (string, error) {
	userLogin, err := u.repository.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid username or Password")
	}

	err = godotenv.Load()
	if err != nil {
		return "", errors.New("invalid username or Password")
	}

	errPassword := helper.CheckPasswordHash(userLogin.Password, user.Password)
	if errPassword != nil {
		return "", errors.New("invalid username or Password")
	}

	token, _ := middleware.GenerateToken(60000, userLogin.UserId, os.Getenv("TOKEN_SECRET"))
	return token, nil
}