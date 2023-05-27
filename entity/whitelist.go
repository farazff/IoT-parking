package entity

import "github.com/google/uuid"

type Whitelist interface {
	Id() int
	PID() uuid.UUID
	CarTag() string
}
