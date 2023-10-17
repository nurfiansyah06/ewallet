package repository

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	TransferAmount(transaction dto.TransactionRequest) (entity.Transaction, error)
}

func NewTransacationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) TransferAmount(transaction dto.TransactionRequest) (entity.Transaction, error) {	
	var transactions entity.Transaction
	
	err := r.db.Create(&transactions).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return entity.Transaction{}, nil
}