package dto

type Wallet struct {
	WalletId   int    `json:"wallet_id"`
	Amount     int    `json:"amount"`
	SourceFund string `json:"source_fund"`
	UserId		int		`json:"user_id"`
	// User       entity.User
}

type WalletResponse struct {
	WalletId     int    `json:"wallet_id"`
	Amount       int    `json:"amount"`
	SourceFund   string `json:"source_fund"`
	WalletNumber int    `json:"wallet_number"`
}