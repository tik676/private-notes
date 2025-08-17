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
				log.Fatal("Error deleting expired notes:", err)
			}
			_, err = DB.Exec(`DELETE FROM refresh_tokens WHERE expires_at < NOW()`)
			if err != nil {
				log.Fatal("Error deleting expired refresh tokens:", err)
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
		if err := res.Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.ExpiresAt, &note.IsPrivate, &note.HashPassword); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Error scanning note:", err)
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
	query := `
        SELECT id, user_id, content, created_at, expires_at, is_private, hash_password
        FROM notes WHERE id=$1 AND is_private = false
    `
	err := DB.QueryRow(query, id).Scan(
		&note.ID,
		&note.UserID,
		&note.Content,
		&note.CreatedAt,
		&note.ExpiresAt,
		&note.IsPrivate,
		&note.HashPassword,
	)

	if err != nil {
		log.Println("GetPublicNote error:", err)
		return nil, err
	}
	return &note, nil
}

func GetNoteByID(noteID int) (*models.Notes, error) {
	var note models.Notes
	query := `SELECT * FROM notes WHERE id=$1`
	err := DB.QueryRow(query, noteID).Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.ExpiresAt, &note.IsPrivate, &note.HashPassword)
	if err != nil {
		return nil, errors.New("Note with this ID does not exist")
	}

	return &note, nil
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

func GetNoteByIDAndUser(noteID, userID int) (*models.Notes, error) {
	var note models.Notes
	query := `SELECT * FROM notes WHERE id=$1 AND user_id=$2`
	err := DB.QueryRow(query, noteID, userID).Scan(&note.ID, &note.UserID, &note.Content, &note.CreatedAt, &note.ExpiresAt, &note.IsPrivate, &note.HashPassword)
	if err != nil {
		return nil, errors.New("No such note")
	}

	return &note, nil
}

func CreateNote(user_id int, content string, duration time.Time, isPrivate bool, hash_Password *string) error {
	query := `INSERT INTO notes(user_id,content,expires_at,is_private,hash_password)VALUES($1, $2, $3, $4, $5);`
	_, err := DB.Exec(query, user_id, content, duration, isPrivate, hash_Password)
	if err != nil {
		return models.ErrToAddNote
	}
	return nil
}

func SaveRefreshToken(userID int, refreshToken string, refreshExp time.Time) error {
	_, err := DB.Exec(`INSERT INTO refresh_tokens(user_id, token, expires_at)VALUES($1, $2, $3);`, userID, refreshToken, refreshExp)
	return err
}

func DeleteRefreshToken(token string) error {
	query := `DELETE FROM refresh_tokens WHERE token=$1`
	_, err := DB.Exec(query, token)
	if err != nil {
		return errors.New("refresh token not received, failed")
	}
	return err
}

func DeleteNote(noteID, userID int) error {
	query := `DELETE FROM notes WHERE id=$1 AND user_id=$2`
	_, err := DB.Exec(query, noteID, userID)
	if err != nil {
		return errors.New("Failed to delete note")
	}
	return err
}

func UpdateNote(noteID, userID int, content string, expiresAt time.Time, isPrivate bool, hashPassword *string) error {
	query := `
        UPDATE notes 
        SET content=$1, expires_at=$2, is_private=$3, hash_password=$4
        WHERE id=$5 AND user_id=$6
    `
	_, err := DB.Exec(query, content, expiresAt, isPrivate, hashPassword, noteID, userID)
	if err != nil {
		return errors.New("Failed to update note")
	}
	return nil
}
