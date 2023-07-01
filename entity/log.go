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

type AdminLog struct {
	ID        int        `json:"id" db:"id"`
	EnterTime time.Time  `json:"enter_time" db:"enter_time"`
	ExitTime  *time.Time `json:"exit_time" db:"exit_time"`
	FirstName string     `json:"first_name" db:"first_name"`
	LastName  string     `json:"last_name" db:"last_name"`
	CarTag    string     `json:"car_tage" db:"car_tag"`
	Phone     string     `json:"phone" db:"phone"`
}
