package entity

import "github.com/google/uuid"

type Whitelist interface {
	Id() int
	PID() uuid.UUID
	CarTag() string
}

type WhitelistDeleteReq struct {
	AdminCode uuid.UUID `json:"admin_code"`
	CarTag    string    `json:"car_tag"`
}

type WhitelistGetReq struct {
	AdminCode uuid.UUID `json:"admin_code"`
}
