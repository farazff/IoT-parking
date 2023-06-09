package entity

type Whitelist interface {
	ID() int
	ParkingID() int
	CarTag() string
}
