package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Parking struct {
	DBID               int       `db:"id"`
	DBName             string    `db:"name"`
	DBAddress          string    `db:"address"`
	DBPhone            string    `db:"phone"`
	DBEnabled          bool      `db:"enabled"`
	DBUuid             uuid.UUID `db:"uuid"`
	DBAccess           int       `db:"access"`
	DBCapacity         int       `db:"capacity"`
	DBRemainedCapacity int       `db:"remained_capacity"`
}

type UserParking struct {
	Parking
	DBAccess int `db:"access"`
}

func (p Parking) ID() int {
	return p.DBID
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

func (p Parking) Uuid() uuid.UUID {
	return p.DBUuid
}

func (uP UserParking) Access() int {
	return uP.DBAccess
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
	DBID        int    `db:"id"`
	DBFirstName string `db:"first_name"`
	DBLastName  string `db:"last_name"`
	DBPhone     string `db:"phone"`
	DBEnabled   bool   `db:"enabled"`
	DBPassword  string `db:"password"`
	DBParkingID int    `db:"parking_id"`
}

func (pa ParkingAdmin) Password() string {
	return pa.DBPassword
}

func (pa ParkingAdmin) ParkingID() int {
	return pa.DBParkingID
}

func (pa ParkingAdmin) ID() int {
	return pa.DBID
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

func (pa ParkingAdmin) Enabled() bool {
	return pa.DBEnabled
}

type Zone struct {
	DBId               int  `db:"id"`
	DBParkingID        int  `db:"parking_id"`
	DBCapacity         int  `db:"capacity"`
	DBRemainedCapacity int  `db:"remained_capacity"`
	DBEnabled          bool `db:"enabled"`
}

func (z Zone) ParkingID() int {
	return z.DBParkingID
}

func (z Zone) ID() int {
	return z.DBId
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

type Whitelist struct {
	DBID        int  `db:"id"`
	DBUserID    int  `db:"user_id"`
	DBParkingID int  `db:"parking_id"`
	DBApproved  bool `db:"approved"`
}

func (w Whitelist) ID() int {
	return w.DBID
}

func (w Whitelist) UserID() int {
	return w.DBUserID
}

func (w Whitelist) ParkingID() int {
	return w.DBParkingID
}

func (w Whitelist) Approved() bool {
	return w.DBApproved
}

type Log struct {
	DBId        int        `db:"id"`
	DBCarTag    string     `db:"car_tag"`
	DBEnterTime time.Time  `db:"enter_time"`
	DBExitTime  *time.Time `db:"exit_time"`
	DBPID       uuid.UUID  `db:"parking_id"`
}

func (l Log) Id() int {
	return l.DBId
}

func (l Log) CarTag() string {
	return l.DBCarTag
}

func (l Log) EnterTime() time.Time {
	return l.DBEnterTime
}

func (l Log) ExitTime() *time.Time {
	return l.DBExitTime
}

func (l Log) ParkingUUID() uuid.UUID {
	return l.DBPID
}

type User struct {
	DBID        int    `db:"ID"`
	DBFirstName string `db:"first_name"`
	DBLastName  string `db:"last_name"`
	DBCarTag    string `db:"car_tag"`
	DBPhone     string `db:"phone"`
	DBPassword  string `db:"password"`
}

func (u User) ID() int {
	return u.DBID
}

func (u User) FirstName() string {
	return u.DBFirstName
}

func (u User) LastName() string {
	return u.DBLastName
}

func (u User) CarTag() string {
	return u.DBCarTag
}

func (u User) Phone() string {
	return u.DBPhone
}

func (u User) Password() string {
	return u.DBPassword
}
