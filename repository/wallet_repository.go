package repository

import "ewalletgolang/entity"

type WalletRepository interface {
	GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error)
}

func (r *repository) GenerateNumberAcoount(wallet entity.Wallet) (entity.Wallet, error) {
	err := r.db.Create(&wallet).Error
	return wallet, err
}