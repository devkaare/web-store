package model

type User struct {
	UserID uint32
	Email  string
	// TODO: Hash this
	Password string
}
