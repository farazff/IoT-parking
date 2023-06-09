package entity

import (
	"github.com/google/uuid"
	"time"
)

type Log interface {
	Id() int
	CarTag() string
	EnterTime() time.Time
	ExitTime() *time.Time
	ParkingUUID() uuid.UUID
}

type CarExit struct {
	ParkingId int    `json:"parking_id"`
	CarTag    string `json:"car_tag"`
}
