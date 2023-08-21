package dto

type WalletRequest struct {
	WalletId   int    `json:"wallet_id"`
	Amount     int    `json:"amount"`
	SourceFund string `json:"source_fund"`
}