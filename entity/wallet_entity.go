package entity

type Wallet struct {
	Wallet_Id    int `json:"wallet_id" gorm:"primaryKey"`
	WalletNumber int `json:"wallet_number"`
	Amount       int `json:"amount"`
	UserId       int `json:"user_id" gorm:"foreignKey:UserRefer"`
	User         User
}