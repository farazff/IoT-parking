package entity

type Whitelist interface {
	ID() int
	UserID() int
	ParkingID() int
	Approved() bool
}

type WhitelistOfficeData struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	CarTag    string `json:"car_tag" db:"car_tag"`
	Phone     string `json:"phone" db:"phone"`
}

type WhitelistUserData struct {
	ID             int    `json:"id" db:"id"`
	ParkingName    string `json:"parking_name" db:"parking_name"`
	ParkingAddress string `json:"parking_address" db:"parking_address"`
	Approved       bool   `json:"approved"`
}
