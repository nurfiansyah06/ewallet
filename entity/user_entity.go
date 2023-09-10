package entity

import "time"

type User struct {
	UserId    int		`json:"user_id" gorm:"primaryKey"`
	Name      string	`json:"name"`
	Email     string	`json:"email"`
	Password	string	`json:"password"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type ProfileUser struct {
	UserId	int	`json:"user_id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}
