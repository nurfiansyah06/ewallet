package repository

import (
	"ewalletgolang/dto"
	"ewalletgolang/entity"
	"fmt"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error)
	TopUpWallet(wallet dto.Wallet) (dto.Wallet, error)
}

func NewWalletRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error) {
	err := r.db.Create(&wallet).Error
	return wallet, err
}

func (r *repository) TopUpWallet(wallet dto.Wallet) (dto.Wallet, error) {	
	var amountWallet entity.Wallet

	r.db.First(&amountWallet, "wallet_id = ?", wallet.WalletId)

	updatedWallet := dto.Wallet{
		WalletId:    wallet.WalletId,
		Amount:      amountWallet.Amount + wallet.Amount,
		SourceFund:  wallet.SourceFund,
	}
	
	result := r.db.Model(&dto.Wallet{}).Where("wallet_id = ?", wallet.WalletId).Updates(updatedWallet)
	if result.Error != nil {
		fmt.Println("Error updating wallet:", result.Error)
		return dto.Wallet{}, result.Error
	}

	return updatedWallet, result.Error
}
