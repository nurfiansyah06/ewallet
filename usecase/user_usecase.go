package usecase

import (
	"errors"
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"
	"ewalletgolang/middleware"
	"ewalletgolang/repository"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Usecase interface {
	Register(user dto.UserRequest) (entity.User, error)
	IsEmailTaken(email string) (bool, error)
	Login(user dto.UserLoginRequest) (string, error)
	ResetPassword(email, newPassword string) error
	FindUserById(userId int) (entity.User, error)
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

	tokenSecret := os.Getenv("TOKEN_SECRET")
	// config, _ := db.LoadConfig(".")
	expirationTime := time.Hour * 24

	token, err := middleware.GenerateToken(expirationTime, userLogin.UserId, tokenSecret)
	log.Println(token, err)
	if err != nil {
    	fmt.Println("Token Generation Error:", err)
    	return "", err	
	}

	return token, nil
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