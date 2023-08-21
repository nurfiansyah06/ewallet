package repository

import (
	"ewalletgolang/entity"

	"gorm.io/gorm"
)

type WalletRepository interface {
	GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error)
	TopUpWallet(wallet entity.Wallet) (entity.Wallet, error)
}

func NewWalletRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error) {
	err := r.db.Create(&wallet).Error
	return wallet, err
}

func (r *repository) TopUpWallet(wallet entity.Wallet) (entity.Wallet, error) {
	var user entity.User
	
	updatedWallet := entity.Wallet{
        Amount:       wallet.Amount,
        SourceFund:   wallet.SourceFund,
    }

    err := r.db.Model(&wallet).Where("user_id = ?", user.UserId).Updates(updatedWallet).Error
    return updatedWallet, err
}