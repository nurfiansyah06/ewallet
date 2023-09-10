package dto

type TransactionRequest struct {
	Amount int    `json:"amount"`
	Wallet Wallet `json:"wallet"`
}