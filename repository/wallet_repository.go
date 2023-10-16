package repository

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository interface {
	TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *walletRepository {
	return &walletRepository{db}
}

func (r *walletRepository) TopUpWallet(wallet dto.Wallet) (entity.WalletResponse, error) {
	var existingWallet entity.Wallet

	if err := r.db.Where("user_id = ?", wallet.UserId).First(&existingWallet).Error; err != nil {
		if err != nil {
			return entity.WalletResponse{}, fmt.Errorf("wallet not found for user with ID")
		}
		return entity.WalletResponse{}, err
	}

	updatedWallet := entity.WalletResponse{
        SourceFund: wallet.SourceFund,
		Amount:     existingWallet.Amount + wallet.Amount,
        WalletNumber: existingWallet.WalletNumber,
	}

	if err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", wallet.UserId).Updates(updatedWallet).Error; err != nil {
		return entity.WalletResponse{}, err
	}

	return updatedWallet, nil
}
