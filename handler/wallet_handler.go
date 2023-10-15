package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type walletHandler struct {
	walletUsecase usecase.WalletUsecase
}

func NewWalletHandler(usecase usecase.WalletUsecase) *walletHandler {
	return &walletHandler{usecase}
}

func (h *walletHandler) TopUpWallet(c *gin.Context) {
	var wallet dto.Wallet

	claimsRaw, _ := c.Get("claims")
	claims, ok := claimsRaw.(jwt.MapClaims)
	userIdRaw, ok := claims["user_id"]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in JWT claims",
		})
		return
	}

	userIdStr := fmt.Sprintf("%v", userIdRaw)
	userIdInt, _ := strconv.Atoi(userIdStr)

	newWallets := dto.Wallet{
		Amount: wallet.Amount,
		SourceFund: wallet.SourceFund,
		UserId: userIdInt,
	}	

	if err := c.ShouldBindJSON(&wallet); err != nil {
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