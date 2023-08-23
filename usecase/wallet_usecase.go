package usecase

import (
	"ewalletgolang/dto"
	"ewalletgolang/repository"
)

type WalletUsecase interface {
	TopUpWallet(wallet dto.Wallet) (dto.Wallet, error)
}

type walletusecase struct {
	repository repository.WalletRepository
}

func NewWalletUsecase(repository repository.WalletRepository) WalletUsecase {
	return &walletusecase{repository}
}

func (u *walletusecase) TopUpWallet(wallet dto.Wallet) (dto.Wallet, error) {
	newWallet, err := u.repository.TopUpWallet(wallet)

	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}
