package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}
