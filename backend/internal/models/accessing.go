package models

type RegisterInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
