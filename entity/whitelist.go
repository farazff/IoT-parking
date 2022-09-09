package entity

type Whitelist interface {
	Id() int
	PID() int
	CarTag() string
}

type WhitelistDeleteReq struct {
	AdminCode int    `json:"admin_code"`
	CarTag    string `json:"car_tag"`
}
