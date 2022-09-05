package entity

import "time"

type SystemAdmin interface {
	Id() int
	FirstName() string
	LastName() string
	Phone() string
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}
