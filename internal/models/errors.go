package models

import "errors"

var (
	ErrInvalidToken    = errors.New("token is invalid")
	ErrExpiredToken    = errors.New("token has expired")
	ErrTypeOfSignature = errors.New("Incorrect type of signature")
)
