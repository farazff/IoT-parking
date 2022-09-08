package entity

type Whitelist interface {
	Id() int
	PID() int
	CarTag() string
}

type WhitelistDeleteReq struct {
	ParkingUuid string `json:"parking_uuid"`
	CarTag      string `json:"car_tag"`
}
