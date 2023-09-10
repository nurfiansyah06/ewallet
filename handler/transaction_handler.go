package handler

import (
	"ewalletgolang/dto"
	"ewalletgolang/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(usecase usecase.TransactionUsecase) *transactionHandler {
	return &transactionHandler{usecase}
}

func (h *transactionHandler) AddAmount(c *gin.Context) {
	var transaction dto.TransactionRequest

	if err := c.ShouldBindJSON(&transaction); err != nil {
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTransaction, err := h.transactionUsecase.AddAmount(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "success",
		"transaction":     newTransaction,
	})
}