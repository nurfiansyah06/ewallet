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
	var wallet dto.Wallet

	if err := c.ShouldBindJSON(&wallet); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	walletId, _ := strconv.Atoi(c.Param("wallet_id"))
	newWallets := dto.Wallet{
		WalletId: walletId,
		Amount: wallet.Amount,
		SourceFund: wallet.SourceFund,
	}

	newWallet, err := h.walletUsecase.TopUpWallet(newWallets)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "success",
		"wallet":    newWallet,
	})

// 	walletId := c.Param("id")

// 	walletIntId, _ := strconv.Atoi(walletId)

// 	walletUpdate := dto.WalletRequest{
// 		WalletId:         walletIntId,
// 		Amount:     newWallet.Amount,
// 		SourceFund: newWallet.SourceFund,
// 	}

// 	walletResponse := dto.WalletResponse{
// 		WalletId: walletIntId,
// 		Amount: newWallet.Amount,
// 		SourceFund: newWallet.SourceFund,
// 	}

// 	wallet, err := h.walletUsecase.TopUpWallet(walletResponse)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"response": "success",
// 		"wallet":   wallet,
// 	})
// }
}