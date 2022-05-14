package models

type User struct {
	Id           int
	Username     string
	Email        string
	PasswordHash string
}
