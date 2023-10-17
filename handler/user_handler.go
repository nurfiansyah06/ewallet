package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/helper"
	"ewalletgolang/middleware"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type userHandler struct {
	userUsecase	usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *userHandler {
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

	loggedUser, err := h.userUsecase.Login(loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	userId, err := h.userUsecase.FindUserById(loggedUser.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}

	token, err := middleware.AuthMiddleware(userId.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
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
	var user entity.User

	userIdRaw, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in context",
		})
		return
	}

	userIdStr := fmt.Sprintf("%v", userIdRaw)
	userIdInt, _ := strconv.Atoi(userIdStr)
	
	user, err := h.userUsecase.FindUserById(userIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"data":     user,
	})
}