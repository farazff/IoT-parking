package entity

type Whitelist interface {
	ID() int
	UserID() int
	ParkingID() int
	Approved() bool
}

type WhitelistOfficeData struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_Name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	CarTag    string `json:"car_tag" db:"car_tag"`
	ParkingID int    `json:"parking_id" db:"parking_id"`
}
