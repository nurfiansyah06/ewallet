package repository

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	AddAmount(transaction dto.TransactionRequest) (entity.Transaction, error)
}

func NewTransacationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddAmount(transaction dto.TransactionRequest) (entity.Transaction, error) {	
	err := r.db.Create(&transaction).Error
	if err != nil {
		return entity.Transaction{}, err
	}

	return entity.Transaction{}, nil
}