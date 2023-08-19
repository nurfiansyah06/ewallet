package entity

type Wallet struct {
	Wallet_Id     int `json:"wallet_id" gorm:"primaryKey"`
	WalletNumber  int `json:"wallet_number"`
	BalanceWallet int `json:"balance_wallet"`
	UserId        int `json:"user_id" gorm:"foreignKey:UserRefer"`
}