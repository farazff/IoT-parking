package rest

import (
	"github.com/farazff/IoT-parking/entity"
	"time"
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
	FUuid      string     `json:"uuid"`
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

func (p Parking) Uuid() string {
	return p.FUuid
}

type ParkingRes struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	Phone     string     `json:"phone"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	Uuid      string     `json:"uuid,omitempty"`
	Capacity  int        `json:"capacity,omitempty"`
}

func toParkingRes(parking entity.Parking, capacity int, id int) ParkingRes {
	response := ParkingRes{
		Id:        parking.Id(),
		Name:      parking.Name(),
		Address:   parking.Address(),
		Phone:     parking.Phone(),
		Enabled:   parking.Enabled(),
		CreatedAt: parking.CreatedAt(),
		UpdatedAt: parking.UpdatedAt(),
		DeletedAt: parking.DeletedAt(),
		Uuid:      parking.Uuid(),
		Capacity:  capacity,
	}
	if id != -1 {
		response.Id = id
	}
	return response
}

func toParkingResSlice(parkings []entity.Parking) []ParkingRes {
	parkingsResSlice := make([]ParkingRes, 0)
	for _, parking := range parkings {
		parkingsResSlice = append(parkingsResSlice, toParkingRes(parking, 0, -1))
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
	FPID       int        `json:"parking_id"`
	FEnabled   bool       `json:"enabled"`
	FCreatedAt time.Time  `json:"createdAt"`
	FUpdatedAt time.Time  `json:"updatedAt"`
	FDeletedAt *time.Time `json:"deletedAt,omitempty"`
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

func (pa ParkingAdmin) PID() int {
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
	Id        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Phone     string     `json:"phone"`
	PID       int        `json:"parking_id"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func toParkingAdminRes(parkingAdmin entity.ParkingAdmin, id int) ParkingAdminRes {
	response := ParkingAdminRes{
		Id:        parkingAdmin.Id(),
		FirstName: parkingAdmin.FirstName(),
		LastName:  parkingAdmin.LastName(),
		PID:       parkingAdmin.PID(),
		Enabled:   parkingAdmin.Enabled(),
		CreatedAt: parkingAdmin.CreatedAt(),
		UpdatedAt: parkingAdmin.UpdatedAt(),
		DeletedAt: parkingAdmin.DeletedAt(),
	}
	if id != -1 {
		response.Id = id
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
	FId               int        `json:"id"`
	FPID              int        `json:"parking_id"`
	FCapacity         int        `json:"capacity"`
	FRemainedCapacity int        `json:"remained_capacity"`
	FEnabled          bool       `json:"enabled"`
	FCreatedAt        time.Time  `json:"created-at"`
	FUpdatedAt        time.Time  `json:"updated-at"`
	FDeletedAt        *time.Time `json:"deleted-at"`
}

func (z Zone) Id() int {
	return z.FId
}

func (z Zone) PID() int {
	return z.FPID
}

func (z Zone) Capacity() int {
	return z.FCapacity
}

func (z Zone) RemainedCapacity() int {
	return z.FRemainedCapacity
}

func (z Zone) Enabled() bool {
	return z.FEnabled
}

func (z Zone) CreatedAt() time.Time {
	return z.FCreatedAt
}

func (z Zone) UpdatedAt() time.Time {
	return z.FUpdatedAt
}

func (z Zone) DeletedAt() *time.Time {
	return z.FDeletedAt
}

type ZoneRes struct {
	Id               int        `json:"id"`
	PID              int        `json:"parking_id"`
	Capacity         int        `json:"capacity"`
	RemainedCapacity int        `json:"remained_capacity"`
	Enabled          bool       `json:"enabled"`
	CreatedAt        time.Time  `json:"created-at"`
	UpdatedAt        time.Time  `json:"updated-at"`
	DeletedAt        *time.Time `json:"deleted-at"`
}

func toZoneRes(zone entity.Zone, id int) ZoneRes {
	response := ZoneRes{
		Id:               zone.Id(),
		PID:              zone.PID(),
		Capacity:         zone.Capacity(),
		RemainedCapacity: zone.RemainedCapacity(),
		Enabled:          zone.Enabled(),
		CreatedAt:        zone.CreatedAt(),
		UpdatedAt:        zone.UpdatedAt(),
		DeletedAt:        zone.DeletedAt(),
	}
	if id != -1 {
		response.Id = id
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
	FId     int    `json:"id"`
	FPID    int    `json:"parking_id"`
	FCarTag string `json:"car_tag"`
}

type WhitelistCreateReq struct {
	Whitelist
	AdminCode int `json:"admin_code"`
}

func (w Whitelist) Id() int {
	return w.FId
}

func (w Whitelist) PID() int {
	return w.FPID
}

func (w Whitelist) CarTag() string {
	return w.FCarTag
}

type WhitelistRes struct {
	Id     int    `json:"id"`
	PID    int    `json:"parking_id"`
	CarTag string `json:"car_tag"`
}

func toWhitelistRes(whitelist entity.Whitelist, id int) WhitelistRes {
	response := WhitelistRes{
		Id:     whitelist.Id(),
		PID:    whitelist.PID(),
		CarTag: whitelist.CarTag(),
	}
	if id != -1 {
		response.Id = id
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
