package error

import "errors"

var (
	ErrInvalidId       = errors.New("invalid id value")
	ErrServer          = errors.New("server failure")
	ErrEmptyFields     = errors.New("empty fields")
	ErrEmptyFullName   = errors.New("empty full name")
	ErrNothingToFound  = errors.New("error: nothing to found")
	ErrNothingToUpdate = errors.New("error: nothing to update")
)
