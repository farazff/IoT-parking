package manager

import "errors"

var (
	ErrParkingNotFound = errors.New("parking does not exist")
	ErrInternalServer  = errors.New("internal server error")
	ErrDuplicateEntity = errors.New("duplicate entity")
)
