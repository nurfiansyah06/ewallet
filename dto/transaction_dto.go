package dto

type TransactionRequest struct {
	From   int    `json:"from"`
	To     int    `json:"to"`
	Amount int    `json:"amount"`
	Wallet Wallet `json:"wallet"`
}