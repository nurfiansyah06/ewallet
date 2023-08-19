package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"

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
			"err": err,
		})
	}

	token, err := h.userUsecase.Login(loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
	}
	
	// response := entity.UserResponse{
	// 	Token: token,
	// }

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"token": token,
	})
}
