package rest

import (
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

type Parking struct {
	FId        int        `json:"id"`
	FName      string     `json:"name"`
	FAddress   string     `json:"address"`
	FPhone     string     `json:"phone"`
	FEnabled   bool       `json:"enabled"`
	FCreatedAt time.Time  `json:"createdAt"`
	FUpdatedAt time.Time  `json:"updatedAt"`
	FDeletedAt *time.Time `json:"deletedAt,omitempty"`
	FUuid      uuid.UUID  `json:"uuid"`
}

func (p Parking) Id() int {
	return p.FId
}

func (p Parking) Name() string {
	return p.FName
}

func (p Parking) Address() string {
	return p.FAddress
}

func (p Parking) Phone() string {
	return p.FPhone
}

func (p Parking) Enabled() bool {
	return p.FEnabled
}

func (p Parking) CreatedAt() time.Time {
	return p.FCreatedAt
}

func (p Parking) UpdatedAt() time.Time {
	return p.FUpdatedAt
}

func (p Parking) DeletedAt() *time.Time {
	return p.FDeletedAt
}

func (p Parking) Uuid() uuid.UUID {
	return p.FUuid
}

type ParkingRes struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Enabled  bool   `json:"enabled"`
	Capacity int    `json:"capacity,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}

func toParkingRes(parking entity.Parking, capacity int, Puuid uuid.UUID) ParkingRes {
	response := ParkingRes{
		Id:       parking.Id(),
		Name:     parking.Name(),
		Address:  parking.Address(),
		Phone:    parking.Phone(),
		Enabled:  parking.Enabled(),
		Capacity: capacity,
	}
	if Puuid != uuid.Nil {
		response.Uuid = Puuid.String()
	}
	return response
}

func toParkingResSlice(parkings []entity.Parking) []ParkingRes {
	parkingsResSlice := make([]ParkingRes, 0)
	for _, parking := range parkings {
		parkingsResSlice = append(parkingsResSlice, toParkingRes(parking, 0, uuid.UUID{}))
	}
	return parkingsResSlice
}

type SystemAdmin struct {
	FId        int        `json:"id"`
	FFirstName string     `json:"first_name"`
	FLastName  string     `json:"last_name"`
	FPhone     string     `json:"phone"`
	FEnabled   bool       `json:"enabled"`
	FCreatedAt time.Time  `json:"createdAt"`
	FUpdatedAt time.Time  `json:"updatedAt"`
	FDeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func (s SystemAdmin) Id() int {
	return s.FId
}

func (s SystemAdmin) FirstName() string {
	return s.FFirstName
}

func (s SystemAdmin) LastName() string {
	return s.FLastName
}

func (s SystemAdmin) Phone() string {
	return s.FPhone
}

func (s SystemAdmin) Enabled() bool {
	return s.FEnabled
}

func (s SystemAdmin) CreatedAt() time.Time {
	return s.FCreatedAt
}

func (s SystemAdmin) UpdatedAt() time.Time {
	return s.FUpdatedAt
}

func (s SystemAdmin) DeletedAt() *time.Time {
	return s.FDeletedAt
}

type SystemAdminRes struct {
	Id        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Phone     string     `json:"phone"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func toSystemAdminRes(SystemAdmin entity.SystemAdmin, id int) SystemAdminRes {
	response := SystemAdminRes{
		Id:        SystemAdmin.Id(),
		FirstName: SystemAdmin.FirstName(),
		LastName:  SystemAdmin.LastName(),
		Enabled:   SystemAdmin.Enabled(),
		CreatedAt: SystemAdmin.CreatedAt(),
		UpdatedAt: SystemAdmin.UpdatedAt(),
		DeletedAt: SystemAdmin.DeletedAt(),
	}
	if id != -1 {
		response.Id = id
	}
	return response
}

func toSystemAdminResSlice(parkings []entity.SystemAdmin) []SystemAdminRes {
	SystemAdminsResSlice := make([]SystemAdminRes, 0)
	for _, SystemAdmin := range parkings {
		SystemAdminsResSlice = append(SystemAdminsResSlice, toSystemAdminRes(SystemAdmin, -1))
	}
	return SystemAdminsResSlice
}

type ParkingAdmin struct {
	FId        int        `json:"id"`
	FFirstName string     `json:"first_name"`
	FLastName  string     `json:"last_name"`
	FPhone     string     `json:"phone"`
	FPID       uuid.UUID  `json:"parking_id"`
	FEnabled   bool       `json:"enabled"`
	FCreatedAt time.Time  `json:"createdAt"`
	FUpdatedAt time.Time  `json:"updatedAt"`
	FDeletedAt *time.Time `json:"deletedAt"`
	FUuid      uuid.UUID  `json:"FUuid"`
}

func (pa ParkingAdmin) Uuid() uuid.UUID {
	return pa.FUuid
}

func (pa ParkingAdmin) Id() int {
	return pa.FId
}

func (pa ParkingAdmin) FirstName() string {
	return pa.FFirstName
}

func (pa ParkingAdmin) LastName() string {
	return pa.FLastName
}

func (pa ParkingAdmin) Phone() string {
	return pa.FPhone
}

func (pa ParkingAdmin) PID() uuid.UUID {
	return pa.FPID
}

func (pa ParkingAdmin) Enabled() bool {
	return pa.FEnabled
}

func (pa ParkingAdmin) CreatedAt() time.Time {
	return pa.FCreatedAt
}

func (pa ParkingAdmin) UpdatedAt() time.Time {
	return pa.FUpdatedAt
}

func (pa ParkingAdmin) DeletedAt() *time.Time {
	return pa.FDeletedAt
}

type ParkingAdminRes struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	PID       uuid.UUID `json:"parking_id"`
	Enabled   bool      `json:"enabled"`
	Uuid      string    `json:"uuid,omitempty"`
}

func toParkingAdminRes(parkingAdmin entity.ParkingAdmin, id int) ParkingAdminRes {
	response := ParkingAdminRes{
		Id:        parkingAdmin.Id(),
		FirstName: parkingAdmin.FirstName(),
		LastName:  parkingAdmin.LastName(),
		Phone:     parkingAdmin.Phone(),
		PID:       parkingAdmin.PID(),
		Enabled:   parkingAdmin.Enabled(),
	}
	if id != -1 {
		response.Id = id
	}
	if parkingAdmin.Uuid() != uuid.Nil {
		response.Uuid = parkingAdmin.Uuid().String()
	}
	return response
}

func toParkingAdminResSlice(parkingAdmins []entity.ParkingAdmin) []ParkingAdminRes {
	ParkingAdminsResSlice := make([]ParkingAdminRes, 0)
	for _, parkingAdmin := range parkingAdmins {
		ParkingAdminsResSlice = append(ParkingAdminsResSlice, toParkingAdminRes(parkingAdmin, -1))
	}
	return ParkingAdminsResSlice
}

type Zone struct {
	FID               int  `json:"id"`
	FCapacity         int  `json:"capacity" validate:"gtefield=FRemainedCapacity"`
	FEnabled          bool `json:"enabled"`
	FRemainedCapacity int  `json:"remained_capacity"`
	FParkingID        int  `json:"parking_id"`
}

func (z Zone) ID() int {
	return z.FID
}

func (z Zone) Capacity() int {
	return z.FCapacity
}

func (z Zone) Enabled() bool {
	return z.FEnabled
}

func (z Zone) RemainedCapacity() int {
	return z.FRemainedCapacity
}

func (z Zone) ParkingID() int {
	return z.FParkingID
}

type ZoneRes struct {
	ID               int  `json:"id"`
	Capacity         int  `json:"capacity"`
	Enabled          bool `json:"enabled"`
	RemainedCapacity int  `json:"remained_capacity"`
	ParkingID        int  `json:"parking_id,omitempty"`
}

func toZoneRes(zone entity.Zone, id int) ZoneRes {
	response := ZoneRes{
		ID:               zone.ID(),
		Capacity:         zone.Capacity(),
		Enabled:          zone.Enabled(),
		RemainedCapacity: zone.RemainedCapacity(),
		ParkingID:        zone.ParkingID(),
	}
	if id != -1 {
		response.ID = id
	}
	return response
}

func toZoneResSlice(zones []entity.Zone) []ZoneRes {
	ZoneResSlice := make([]ZoneRes, 0)
	for _, zone := range zones {
		ZoneResSlice = append(ZoneResSlice, toZoneRes(zone, -1))
	}
	return ZoneResSlice
}

type Whitelist struct {
	FID        int    `json:"id"`
	FParkingID int    `json:"parking_id validate:"required"`
	FCarTag    string `json:"car_tag" validate:"required"`
}

func (w Whitelist) ID() int {
	return w.FID
}

func (w Whitelist) ParkingID() int {
	return w.FParkingID
}

func (w Whitelist) CarTag() string {
	return w.FCarTag
}

type WhitelistRes struct {
	ID        int    `json:"id"`
	ParkingID int    `json:"parking_id"`
	CarTag    string `json:"car_tag"`
}

func toWhitelistRes(whitelist entity.Whitelist, id int) WhitelistRes {
	response := WhitelistRes{
		ID:        whitelist.ID(),
		ParkingID: whitelist.ParkingID(),
		CarTag:    whitelist.CarTag(),
	}
	if id != -1 {
		response.ID = id
	}
	return response
}

func toWhitelistResSlice(whitelists []entity.Whitelist) []WhitelistRes {
	whitelistResSlice := make([]WhitelistRes, 0)
	for _, whitelist := range whitelists {
		whitelistResSlice = append(whitelistResSlice, toWhitelistRes(whitelist, -1))
	}
	return whitelistResSlice
}

type Log struct {
	FId        int        `json:"id"`
	FCarTag    string     `json:"car_tag"`
	FEnterTime time.Time  `json:"enter_time"`
	FExitTime  *time.Time `json:"exit_time,omitempty"`
	FPID       uuid.UUID  `json:"parking_id"`
}

func (l Log) Id() int {
	return l.FId
}

func (l Log) CarTag() string {
	return l.FCarTag
}

func (l Log) EnterTime() time.Time {
	return l.FEnterTime
}

func (l Log) ExitTime() *time.Time {
	return l.FExitTime
}

func (l Log) ParkingUUID() uuid.UUID {
	return l.FPID
}

type LogRes struct {
	Id        int        `json:"id"`
	CarTag    string     `json:"car_tag"`
	EnterTime time.Time  `json:"enter_time"`
	ExitTime  *time.Time `json:"exit_time,omitempty"`
	PID       uuid.UUID  `json:"parking_id"`
}

func toLogRes(log entity.Log, id int) LogRes {
	response := LogRes{
		Id:        log.Id(),
		CarTag:    log.CarTag(),
		EnterTime: log.EnterTime(),
		ExitTime:  log.ExitTime(),
		PID:       log.ParkingUUID(),
	}
	if id != -1 {
		response.Id = id
	}
	return response
}

func toLogResSlice(logs []entity.Log) []LogRes {
	logResSlice := make([]LogRes, 0)
	for _, log := range logs {
		logResSlice = append(logResSlice, toLogRes(log, -1))
	}
	return logResSlice
}
