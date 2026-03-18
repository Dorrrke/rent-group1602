package errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

const (
	IncorrectFieldValuesError = "Incorrect field values: %w"
)
