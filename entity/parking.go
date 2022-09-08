package entity

import "time"

type Parking interface {
	Id() int
	Name() string
	Address() string
	Phone() string
	Enabled() bool
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
	Uuid() string
}
