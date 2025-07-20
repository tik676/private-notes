package models

import "errors"

var (
	// token errors
	ErrInvalidToken    = errors.New("Token is invalid")
	ErrExpiredToken    = errors.New("Token has expired")
	ErrTypeOfSignature = errors.New("Incorrect type of signature")

	// user errors
	ErrUserNotFound = errors.New("User not found")

	// db errors
	ErrToAddNote = errors.New("Failed to add note")
)
