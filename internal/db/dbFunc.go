package db

import (
	"database/sql"
	"log"
	"private-notes/internal/models"
	"time"
)

var (
	now = time.Now()
)

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

func CreateNote(user_id int, content string, duration time.Time, isPrivate bool) error {
	query := `INSERT INTO notes(user_id,content,expires_at,is_private)VALUES($1, $2, $3,$4);`
	_, err := DB.Exec(query, user_id, content, duration, isPrivate)
	if err != nil {
		return models.ErrToAddNote
	}
	return nil
}
