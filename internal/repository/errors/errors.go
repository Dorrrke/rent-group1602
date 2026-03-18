package errors

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User with this email already exists")
	ErrUserNotFound      = errors.New("User with this email not fount")
)

var (
	ErrCarAlreadyExists  = errors.New("This car is already registered")
	ErrNotAvailableCars  = errors.New("Now we don't have available cars")
	ErrRentAlreadyExists = errors.New("This car is already rented")
	ErrRentNotFound      = errors.New("Not found rent with this id")
)
