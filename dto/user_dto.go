package dto

type UserRequest struct {
	Name     string `json:"name"`
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