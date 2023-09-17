package dto

type UserRequest struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResetPasssword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}