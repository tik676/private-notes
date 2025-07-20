package authorization

import (
	"log"
	"private-notes/internal/db"
)

func RegisterUser(name, password string) error {
	PasswordHash, err := GenerateHash(password)
	if err != nil {
		return err
	}

	query := `INSERT INTO users(name,password_hash)VALUES ($1, $2)`
	_, err = db.DB.Exec(query, name, PasswordHash)
	if err != nil {
		log.Println("❌ Ошибка при создании пользователя:", err)
		return err
	}

	return nil
}
