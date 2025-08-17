package authorization

import (
	"database/sql"
	"errors"
	"private-notes/internal/db"
	"private-notes/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func LoginUser(name, password string) (int, error) {
	var user models.User
	query := `SELECT * FROM Users WHERE name = $1`
	err := db.DB.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("error user not found")
		}
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return 0, errors.New("wrong password")
	}

	return user.ID, nil
}
