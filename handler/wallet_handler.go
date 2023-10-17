package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type walletHandler struct {
	walletUsecase usecase.WalletUsecase
	userUsecase		usecase.UserUsecase
}

func NewWalletHandler(usecase usecase.WalletUsecase, userUsecase usecase.UserUsecase) *walletHandler {
	return &walletHandler{
		walletUsecase: usecase,
		userUsecase:   userUsecase,
	}
}

func (h *walletHandler) TopUpWallet(c *gin.Context) {
	var wallet dto.Wallet

	userIdInt, _ := c.Get("user_id")
	userId, _ := userIdInt.(int)

	_, err := h.userUsecase.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})

		return
	}

	newWallets := dto.Wallet{
		Balance: wallet.Balance,
		SourceFund: wallet.SourceFund,
		UserId: userId,
	}	

	if err := c.ShouldBindJSON(&newWallets); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	newWallet, err := h.walletUsecase.TopUpWallet(newWallets)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"wallet":    newWallet,
	})
}