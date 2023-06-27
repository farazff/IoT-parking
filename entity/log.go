package entity

import (
	"time"
)

type Log interface {
	ID() int
	UserID() int
	EnterTime() time.Time
	ExitTime() *time.Time
	ParkingID() int
}

type UserLog struct {
	ID             int        `json:"id" db:"id"`
	EnterTime      time.Time  `json:"enter_time" db:"enter_time"`
	ExitTime       *time.Time `json:"exit_time" db:"exit_time"`
	ParkingName    string     `json:"parking_name" db:"parking_name"`
	ParkingAddress string     `json:"parking_address" db:"parking_address"`
}
