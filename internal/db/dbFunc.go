package db

import (
	"database/sql"
	"errors"
	"log"
	"private-notes/internal/models"
	"time"
)

func RegularClearNoteByExpires() {
	go func() {
		for {
			_, err := DB.Exec(`DELETE FROM notes WHERE expires_at < NOW()`)
			if err != nil {
				log.Fatal("Ошибка удаления просроченных заметок:", err)
			}
			_, err = DB.Exec(`DELETE FROM refresh_tokens WHERE expires_at < NOW()`)
			if err != nil {
				log.Fatal("Ошибка удаления просроченных refresh токенов:", err)
			}
			time.Sleep(1 * time.Minute)
		}
	}()
}

func GetWithIDNotesMe(userID int) ([]models.Notes, error) {
	query := `SELECT * FROM notes WHERE user_id=$1`
	res, err := DB.Query(query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrUserNotFound
		}
		return nil, err
	}

	var userNotes []models.Notes

	for res.Next() {
		var note models.Notes

		if err := res.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.ExpiresAt, &note.IsPrivate); err != nil {
			if err == sql.ErrNoRows {
				log.Println("ошибка сканирования заметки:", err)
				return nil, models.ErrUserNotFound
			}
			return nil, err
		}

		userNotes = append(userNotes, note)
	}
	return userNotes, nil

}

func GetPublicNote(id int) (*models.Notes, error) {
	var note models.Notes
	query := `SELECT * FROM notes WHERE id=$1 AND is_private = false`
	err := DB.QueryRow(query, id).Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.ExpiresAt, &note.IsPrivate)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func CreateNote(user_id int, content string, duration time.Time, isPrivate bool) error {
	query := `INSERT INTO notes(user_id,content,expires_at,is_private)VALUES($1, $2, $3,$4);`
	_, err := DB.Exec(query, user_id, content, duration, isPrivate)
	if err != nil {
		return models.ErrToAddNote
	}
	return nil
}

func SaveRefreshToken(userID int, refreshToken string, refreshExp time.Time) error {
	_, err := DB.Exec(`INSERT INTO refresh_tokens(user_id, token, expires_at)VALUES($1, $2, $3);`, userID, refreshToken, refreshExp)
	return err
}

func GetUserIDByRefreshToken(token string) (int, error) {
	var userID int
	var expiresAt time.Time

	query := `SELECT user_id, expires_at FROM refresh_tokens WHERE token = $1`
	err := DB.QueryRow(query, token).Scan(&userID, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("refresh token not found")
		}
		return 0, err
	}

	if time.Now().After(expiresAt) {
		return 0, errors.New("refresh token expired")
	}

	return userID, nil
}

func DeleteRefreshToken(token string) error {
	query := `DELETE * FROM refresh_tokens WHERE token=$1`
	_, err := DB.Exec(query, token)
	if err != nil {
		return errors.New("refresh token ne prishel kaput")
	}
	return err
}
