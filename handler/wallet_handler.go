package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type walletHandler struct {
	walletUsecase usecase.WalletUsecase
}

func NewWalletHandler(usecase usecase.WalletUsecase) *walletHandler {
	return &walletHandler{usecase}
}

func (h *walletHandler) TopUpWallet(c *gin.Context) {
	var newWallet dto.WalletRequest

	if err := c.ShouldBindJSON(&newWallet); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	walletId := c.Param("id")

	walletIntId, _ := strconv.Atoi(walletId)

	walletUpdate := dto.WalletRequest{
		WalletId:         walletIntId,
		Amount:     newWallet.Amount,
		SourceFund: newWallet.SourceFund,
	}

	wallet, err := h.walletUsecase.TopUpWallet(walletUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"wallet":   wallet,
	})
}
