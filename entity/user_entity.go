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

type UserResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}