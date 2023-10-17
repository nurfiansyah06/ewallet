package usecase

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/repository"
)

type WalletUsecase interface {
	TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error)
	FindUserById(userId int) (entity.WalletResponse, error)
}

type walletusecase struct {
	repository repository.WalletRepository
}

func NewWalletUsecase(repository repository.WalletRepository) WalletUsecase {
	return &walletusecase{repository}
}

func (u *walletusecase) TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error) {
	newWallet, err := u.repository.TopUpWallet(wallet)

	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}

func (u *walletusecase) FindUserById(userId int) (entity.WalletResponse, error)  {
	user, err := u.repository.FindWalletByUserId(userId)
	if err != nil {
		return entity.WalletResponse{}, err
	}

	return user, nil
}