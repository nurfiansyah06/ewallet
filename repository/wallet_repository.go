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

func (r *repository) TopUpWallet(wallet dto.Wallet) (entity.Wallet, error) {	
	var existingWallet entity.Wallet

    // Retrieve the wallet record based on user_id
    if err := r.db.Where("user_id = ?", wallet.UserId).First(&existingWallet).Error; err != nil {
        // Handle the error if the wallet record is not found
        if err != nil {
            return entity.Wallet{}, fmt.Errorf("wallet not found for user with ID")
        }
        // Handle other database errors
        return entity.Wallet{}, err
    }

    // Calculate the updated amount by adding the new amount to the existing amount
    updatedAmount := existingWallet.Amount + wallet.Amount

    // Create an updated wallet DTO with the new amount
    updatedWallet := entity.Wallet{
        WalletId:    wallet.WalletId,
        Amount:      updatedAmount,
        SourceFund:  wallet.SourceFund,
        UserId: wallet.UserId, // Make sure to include the user ID
    }

    // Update the wallet record in the database
    if err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", wallet.UserId).Updates(updatedWallet).Error; err != nil {
        // Handle the error if the update fails
        return entity.Wallet{}, err
    }

    // Return the updated wallet DTO and nil error to indicate success
    return updatedWallet, nil
}
