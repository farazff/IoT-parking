package manager

import "errors"

var (
	ErrNotFound        = errors.New("entity does not exist")
	ErrInternalServer  = errors.New("internal server error")
	ErrDuplicateEntity = errors.New("duplicate entity")
)