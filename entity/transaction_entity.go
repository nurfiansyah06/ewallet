package entity

import "time"

type Transaction struct {
	TransactionId int	`json:"transaction_id"`
	From		int		`json:"from"`
	To			int		`json:"to"`
	Amount        int	`json:"amount"`
	// Wallet        Wallet	`json:"wallet"`
	CreatedAt     time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	// WalletId	int	`json:"wallet_id" gorm:"foreignKey:WalletRefer"`
}