package app

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrCountryNotFound   = errors.New("country not found")
	ErrEmptyCountryCode  = errors.New("country code cannot be empty")
	ErrEmptyGivenRole    = errors.New("role cannot be empty")
)
