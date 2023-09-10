package usecase

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/repository"
)

type TransactionUsecase interface {
	AddAmount(transaction dto.TransactionRequest) (entity.Transaction, error)
}

type transactionusecase struct {
	repository repository.TransactionRepository
}

func NewTransactionUsecase(repository repository.TransactionRepository) TransactionUsecase {
	return &transactionusecase{repository}
}

func (u *transactionusecase) AddAmount(transaction dto.TransactionRequest) (entity.Transaction, error) {
	newTransaction, err := u.repository.AddAmount(transaction)

	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}