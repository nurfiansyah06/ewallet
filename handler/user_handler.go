package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/helper"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type userHandler struct {
	userUsecase	usecase.Usecase
}

func NewUserHandler(usecase usecase.Usecase) *userHandler {
	return &userHandler{usecase}
}

func (h *userHandler) Register(c *gin.Context) {
	var user dto.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailTaken, _ := h.userUsecase.IsEmailTaken(user.Email)

	if emailTaken {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already taken",
		})
		return
	}

	newUser, err := h.userUsecase.Register(user)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s", e.Field())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"user":     newUser,
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var loginRequest dto.UserLoginRequest
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "cant binding",
		})
		return
	}

	token, err := h.userUsecase.Login(loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email or password is not correct",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"token": token,
	})
}


func (h *userHandler) ResetPassword(c *gin.Context)  {
	var reset dto.UserResetPasssword

	err := c.ShouldBindJSON(&reset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	pass,_ := helper.HashPassword(reset.Password)

	err = h.userUsecase.ResetPassword(reset.Email, pass)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Password reset successfully"})
}

func (h *userHandler) FindUserById(c *gin.Context) {
	// var user entity.User
	userId := c.Param("id")

	intUser, _ := strconv.Atoi(userId)
	user, err := h.userUsecase.FindUserById(intUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
		
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"data": user,
	})
}