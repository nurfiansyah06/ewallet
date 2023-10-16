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

	userId, err := h.userUsecase.FindUserById(userIdInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})

		return
	}

	// fmt.Println(wallet.Amount)
	fmt.Println(userId.UserId)
	newWallets := dto.Wallet{
		Amount: wallet.Amount,
		SourceFund: wallet.SourceFund,
		UserId: userId.UserId,
	}	

	fmt.Println(newWallets.Amount)
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