package repository

import "errors"

var (
	// ErrNotFound represent no row exist on database
	ErrNotFound = errors.New("not found")

	ErrDuplicateEntity = errors.New("duplicate")
)
