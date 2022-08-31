package postgres

import "time"

type Parking struct {
	DBId        int        `db:"id"`
	DBName      string     `db:"name"`
	DBAddress   string     `db:"address"`
	DBPhone     string     `db:"phone"`
	DBEnabled   bool       `db:"enabled"`
	DBCreatedAt time.Time  `db:"created_at"`
	DBUpdatedAt time.Time  `db:"updated_at"`
	DBDeletedAt *time.Time `db:"deleted_at"`
}

func (p Parking) Id() int {
	return p.DBId
}

func (p Parking) Name() string {
	return p.DBName
}

func (p Parking) Address() string {
	return p.DBAddress
}

func (p Parking) Phone() string {
	return p.DBPhone
}

func (p Parking) Enabled() bool {
	return p.DBEnabled
}

func (p Parking) CreatedAt() time.Time {
	return p.DBCreatedAt
}

func (p Parking) UpdatedAt() time.Time {
	return p.DBUpdatedAt
}

func (p Parking) DeletedAt() *time.Time {
	return p.DBDeletedAt
}
