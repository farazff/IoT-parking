package entity

import (
	"github.com/google/uuid"
)

type Parking interface {
	ID() int
	Name() string
	Address() string
	Phone() string
	Enabled() bool
	Uuid() uuid.UUID
}
