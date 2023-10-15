package usecase

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"ewalletgolang/repository"
)

type WalletUsecase interface {
	TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error)
}

type walletusecase struct {
	repository repository.WalletRepository
}

func NewWalletUsecase(repository repository.WalletRepository) WalletUsecase {
	return &walletusecase{repository}
}

func (u *walletusecase) TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error) {
	updatedWallet := dto.Wallet{
		Amount:     wallet.Amount,
		SourceFund: wallet.SourceFund,
	}

	newWallet, err := u.repository.TopUpWallet(updatedWallet)

	if err != nil {
		return newWallet, err
	}

	return newWallet, nil
}