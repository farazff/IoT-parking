package entity

import (
	"github.com/google/uuid"
	"time"
)

type Zone interface {
	Id() int
	PID() uuid.UUID
	Capacity() int
	RemainedCapacity() int
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}
