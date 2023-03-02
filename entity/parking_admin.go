package entity

import (
	"github.com/google/uuid"
	"time"
)

type ParkingAdmin interface {
	Id() int
	FirstName() string
	LastName() string
	Phone() string
	PID() uuid.UUID
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}
