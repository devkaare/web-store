package model

type User struct {
	UserID   uint32 `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
