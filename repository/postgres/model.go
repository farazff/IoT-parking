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

type SystemAdmin struct {
	DBId        int        `db:"id"`
	DBFirstName string     `db:"first_name"`
	DBLastName  string     `db:"last_name"`
	DBPhone     string     `db:"phone"`
	DBEnabled   bool       `db:"enabled"`
	DBCreatedAt time.Time  `db:"createdAt"`
	DBUpdatedAt time.Time  `db:"updatedAt"`
	DBDeletedAt *time.Time `db:"deletedAt,omitempty"`
}

func (s SystemAdmin) Id() int {
	return s.DBId
}

func (s SystemAdmin) FirstName() string {
	return s.DBFirstName
}

func (s SystemAdmin) LastName() string {
	return s.DBLastName
}

func (s SystemAdmin) Phone() string {
	return s.DBPhone
}

func (s SystemAdmin) Enabled() bool {
	return s.DBEnabled
}

func (s SystemAdmin) CreatedAt() time.Time {
	return s.DBCreatedAt
}

func (s SystemAdmin) UpdatedAt() time.Time {
	return s.DBUpdatedAt
}

func (s SystemAdmin) DeletedAt() *time.Time {
	return s.DBDeletedAt
}
