package entity

type ParkingAdmin interface {
	ID() int
	FirstName() string
	LastName() string
	Phone() string
	Enabled() bool
	Password() string
	ParkingID() int
}
