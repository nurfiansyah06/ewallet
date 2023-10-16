package dto

type Wallet struct {
	Amount     int    `json:"amount"`
	SourceFund string `json:"source_fund"`
	UserId     int    `json:"user_id"`
}
