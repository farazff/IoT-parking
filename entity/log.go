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

type CarExit struct {
	ParkingId int    `json:"parking_id"`
	CarTag    string `json:"car_tag"`
}
