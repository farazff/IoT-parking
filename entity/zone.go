package entity

import "time"

type Zone interface {
	Id() int
	PID() int
	ZID() int
	Capacity() int
	RemainedCapacity() int
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}
