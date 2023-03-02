package entity

import "github.com/google/uuid"

type Whitelist interface {
	Id() int
	PID() uuid.UUID
	CarTag() string
}

type WhitelistDeleteReq struct {
	AdminCode int    `json:"admin_code"`
	CarTag    string `json:"car_tag"`
}

type WhitelistGetReq struct {
	AdminCode int `json:"admin_code"`
}
