package entity

import "time"

type Transaction struct {
	TransactionId int	`json:"transaction_id"`
	Amount        int	`json:"amount"`
	Wallet        Wallet	`json:"wallet"`
	CreatedAt     time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}