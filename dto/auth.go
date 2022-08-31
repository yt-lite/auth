package dto

type RegisterAuthRequest struct {
	Email string `json:"email"`
}

type LoginAuthRequest struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
