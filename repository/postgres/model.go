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

func (sa SystemAdmin) Id() int {
	return sa.DBId
}

func (sa SystemAdmin) FirstName() string {
	return sa.DBFirstName
}

func (sa SystemAdmin) LastName() string {
	return sa.DBLastName
}

func (sa SystemAdmin) Phone() string {
	return sa.DBPhone
}

func (sa SystemAdmin) Enabled() bool {
	return sa.DBEnabled
}

func (sa SystemAdmin) CreatedAt() time.Time {
	return sa.DBCreatedAt
}

func (sa SystemAdmin) UpdatedAt() time.Time {
	return sa.DBUpdatedAt
}

func (sa SystemAdmin) DeletedAt() *time.Time {
	return sa.DBDeletedAt
}

type ParkingAdmin struct {
	DBId        int        `db:"id"`
	DBFirstName string     `db:"first_name"`
	DBLastName  string     `db:"last_name"`
	DBPhone     string     `db:"phone"`
	DBPID       int        `db:"p_id"`
	DBEnabled   bool       `db:"enabled"`
	DBCreatedAt time.Time  `db:"createdAt"`
	DBUpdatedAt time.Time  `db:"updatedAt"`
	DBDeletedAt *time.Time `db:"deletedAt,omitempty"`
}

func (pa ParkingAdmin) Id() int {
	return pa.DBId
}

func (pa ParkingAdmin) FirstName() string {
	return pa.DBFirstName
}

func (pa ParkingAdmin) LastName() string {
	return pa.DBLastName
}

func (pa ParkingAdmin) Phone() string {
	return pa.DBPhone
}

func (pa ParkingAdmin) PID() int {
	return pa.DBPID
}

func (pa ParkingAdmin) Enabled() bool {
	return pa.DBEnabled
}

func (pa ParkingAdmin) CreatedAt() time.Time {
	return pa.DBCreatedAt
}

func (pa ParkingAdmin) UpdatedAt() time.Time {
	return pa.DBUpdatedAt
}

func (pa ParkingAdmin) DeletedAt() *time.Time {
	return pa.DBDeletedAt
}

type Zone struct {
	DBId               int        `db:"id"`
	DBPID              int        `db:"p_id"`
	DBCapacity         int        `db:"capacity"`
	DBRemainedCapacity int        `json:"remained_capacity"`
	DBEnabled          bool       `db:"enabled"`
	DBCreatedAt        time.Time  `db:"created-at"`
	DBUpdatedAt        time.Time  `db:"updated-at"`
	DBDeletedAt        *time.Time `db:"deleted-at"`
}

func (z Zone) Id() int {
	return z.DBId
}

func (z Zone) PID() int {
	return z.DBPID
}

func (z Zone) Capacity() int {
	return z.DBCapacity
}

func (z Zone) RemainedCapacity() int {
	return z.DBRemainedCapacity
}

func (z Zone) Enabled() bool {
	return z.DBEnabled
}

func (z Zone) CreatedAt() time.Time {
	return z.DBCreatedAt
}

func (z Zone) UpdatedAt() time.Time {
	return z.DBUpdatedAt
}

func (z Zone) DeletedAt() *time.Time {
	return z.DBDeletedAt
}
