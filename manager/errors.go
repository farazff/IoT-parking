package manager

import "errors"

var (
	ErrNotFound        = errors.New("entity does not exist")
	ErrInternalServer  = errors.New("internal server error")
	ErrDuplicateEntity = errors.New("duplicate entity")
	ErrNoAccess        = errors.New("no access")
	ErrInvalidCarTag   = errors.New("invalid car tag")
	ErrParkingNotFound = errors.New("parking not found")
)
