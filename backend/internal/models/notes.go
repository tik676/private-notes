package models

import "time"

type Notes struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
	IsPrivate    bool      `json:"is_private"`
	HashPassword *string   `json:"-"`
}

type NoteInput struct {
	Content   string    `json:"content"`
	ExpiresAt time.Time `json:"expires_at"`
	IsPrivate bool      `json:"is_private"`
	Password  *string   `json:"password,omitempty"`
}

type UpdateNote struct {
	Content   *string    `json:"content"`
	ExpiresAt *time.Time `json:"expires_at"`
	IsPrivate *bool      `json:"is_private"`
	Password  *string    `json:"password,omitempty"`
}
