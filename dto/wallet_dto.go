package dto

type Wallet struct {
	Balance    int    `json:"balance"`
	SourceFund string `json:"source_fund"`
	UserId     int    `json:"user_id"`
}
