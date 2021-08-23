package models

type User struct {
	UserID          int    `json:"user_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
}
