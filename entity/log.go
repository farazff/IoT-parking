package entity

import (
	"time"
)

type Log interface {
	ID() int
	CarTag() string
	EnterTime() time.Time
	ExitTime() *time.Time
	ParkingID() int
}
