package errors

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid credentials")
)

const (
	IncorrectFieldValuesError = "Incorrect field values: %w"
)

// Cars errors.
var (
	ErrCarNotAvailable = errors.New("Car is not available")
)
