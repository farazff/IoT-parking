package entity

import "time"

type ParkingAdmin interface {
	Id() int
	FirstName() string
	LastName() string
	Phone() string
	PID() int
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}
