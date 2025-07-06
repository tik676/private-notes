package db

import (
	"database/sql"
	"private-notes/internal/models"
)

func GetWithIDNotesMe(userID int) ([]models.Notes, error) {
	query := `SELECT * FROM posts WHERE user_id=$1`
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
				return nil, models.ErrUserNotFound
			}
			return nil, err
		}

		userNotes = append(userNotes, note)
	}
	return userNotes, nil

}
