package app

import "errors"

var (
	ErrUserAlreadyExists                 = errors.New("user already exists")
	ErrCountryNotFound                   = errors.New("country not found")
	ErrEmptyCountryCode                  = errors.New("country code cannot be empty")
	ErrEmptyGivenRole                    = errors.New("role cannot be empty")
	ErrOneOfRequiredFieldsEmpty          = errors.New("all fields are required")
	ErrCountryWithGivenCodeAlreadyExists = errors.New("country with given code already exists")
	ErrTripWithGivenCodeAlreadyExists    = errors.New("trip with given code already exists")
	ErrTripWithGivenCallbackNotFound     = errors.New("trip with given callback not found")
	ErrContentNotFound                   = errors.New("content not found")
)
