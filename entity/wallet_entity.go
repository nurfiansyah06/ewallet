package entity

type Wallet struct {
	WalletId     int    `json:"wallet_id" gorm:"primaryKey"`
	WalletNumber int    `json:"wallet_number"`
	Amount       int    `json:"amount"`
	SourceFund   string `json:"source_fund"`
	UserId       int    `json:"user_id" gorm:"foreignKey:UserRefer"`
}