package repository

import "errors"

var (
	// ErrNotFound represent no row exist on database
	ErrNotFound                    = errors.New("not found")
	ErrParkingForeignKeyConstraint = errors.New("parking does not exist")
	ErrDuplicateEntity             = errors.New("duplicate")
)
