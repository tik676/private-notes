package models

import "errors"

var (
	// token errors
	ErrInvalidToken    = errors.New("token is invalid")
	ErrExpiredToken    = errors.New("token has expired")
	ErrTypeOfSignature = errors.New("Incorrect type of signature")

	// user errors
	ErrUserNotFound = errors.New("User not found")
)
