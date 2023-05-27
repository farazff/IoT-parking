package entity

type Zone interface {
	ID() int
	Capacity() int
	Enabled() bool
	RemainedCapacity() int
	ParkingID() int
}
