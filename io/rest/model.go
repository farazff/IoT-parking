package rest

import (
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

type Parking struct {
	FID      int       `json:"id"`
	FName    string    `json:"name" validate:"required"`
	FAddress string    `json:"address" validate:"required"`
	FPhone   string    `json:"phone" validate:"required"`
	FEnabled bool      `json:"enabled"`
	FUuid    uuid.UUID `json:"uuid"`
}

func (p Parking) ID() int {
	return p.FID
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

func (p Parking) Uuid() uuid.UUID {
	return p.FUuid
}

type ParkingRes struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Enabled  bool   `json:"enabled"`
	Capacity int    `json:"capacity,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}

func toParkingRes(parking entity.Parking, capacity int, Puuid uuid.UUID) ParkingRes {
	response := ParkingRes{
		ID:       parking.ID(),
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
	FID        int    `json:"id"`
	FFirstName string `json:"first_name"`
	FLastName  string `json:"last_name"`
	FPhone     string `json:"phone"`
	FEnabled   bool   `json:"enabled"`
	FPassword  string `json:"password"`
	FParkingID int    `json:"parking_id"`
}

func (pa ParkingAdmin) ID() int {
	return pa.FID
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

func (pa ParkingAdmin) Enabled() bool {
	return pa.FEnabled
}

func (pa ParkingAdmin) Password() string {
	return pa.FPassword
}

func (pa ParkingAdmin) ParkingID() int {
	return pa.FParkingID
}

type ParkingAdminRes struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Enabled   bool   `json:"enabled"`
	Password  string `json:"password"`
	ParkingID int    `json:"parking_id"`
}

func toParkingAdminRes(parkingAdmin entity.ParkingAdmin, id int) ParkingAdminRes {
	response := ParkingAdminRes{
		ID:        parkingAdmin.ID(),
		FirstName: parkingAdmin.FirstName(),
		LastName:  parkingAdmin.LastName(),
		Phone:     parkingAdmin.Phone(),
		ParkingID: parkingAdmin.ParkingID(),
		Enabled:   parkingAdmin.Enabled(),
	}
	if id != -1 {
		response.ID = id
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
	FID        int  `json:"id"`
	FUserID    int  `json:"user_id"`
	FParkingID int  `json:"parking_id,validate:required"`
	FApproved  bool `json:"approved"`
}

func (w Whitelist) ID() int {
	return w.FID
}

func (w Whitelist) UserID() int {
	return w.FUserID
}

func (w Whitelist) ParkingID() int {
	return w.FParkingID
}

func (w Whitelist) Approved() bool {
	return w.FApproved
}

func toWhitelistOfficeRes(whitelist entity.WhitelistOfficeData) entity.WhitelistOfficeData {
	response := entity.WhitelistOfficeData{
		ID:        whitelist.ID,
		FirstName: whitelist.FirstName,
		LastName:  whitelist.LastName,
		CarTag:    whitelist.CarTag,
		ParkingID: whitelist.ParkingID,
	}
	return response
}

func toWhitelistOfficeResSlice(whitelists []entity.WhitelistOfficeData) []entity.WhitelistOfficeData {
	whitelistResSlice := make([]entity.WhitelistOfficeData, 0)
	for _, whitelist := range whitelists {
		whitelistResSlice = append(whitelistResSlice, toWhitelistOfficeRes(whitelist))
	}
	return whitelistResSlice
}

func toWhitelistUserRes(whitelist entity.WhitelistUserData) entity.WhitelistUserData {
	response := entity.WhitelistUserData{
		ID:             whitelist.ID,
		ParkingName:    whitelist.ParkingName,
		ParkingAddress: whitelist.ParkingAddress,
		Approved:       whitelist.Approved,
	}
	return response
}

func toWhitelistUserResSlice(whitelists []entity.WhitelistUserData) []entity.WhitelistUserData {
	whitelistResSlice := make([]entity.WhitelistUserData, 0)
	for _, whitelist := range whitelists {
		whitelistResSlice = append(whitelistResSlice, toWhitelistUserRes(whitelist))
	}
	return whitelistResSlice
}

type Log struct {
	FID        int        `json:"id"`
	FCarTag    string     `json:"car_tag,validate:required"`
	FEnterTime time.Time  `json:"enter_time"`
	FExitTime  *time.Time `json:"exit_time,omitempty"`
	FParkingID int        `json:"parking_id"`
}

func (l Log) ID() int {
	return l.FID
}

func (l Log) UserID() int {
	return 0
}

func (l Log) EnterTime() time.Time {
	return l.FEnterTime
}

func (l Log) ExitTime() *time.Time {
	return l.FExitTime
}

func (l Log) ParkingID() int {
	return l.FParkingID
}

type LogRes struct {
	ID        int        `json:"id"`
	CarTag    string     `json:"car_tag"`
	EnterTime time.Time  `json:"enter_time"`
	ExitTime  *time.Time `json:"exit_time,omitempty"`
	ParkingID int        `json:"parking_id"`
}

func toUserLogsRes(whitelist entity.UserLog) entity.UserLog {
	response := entity.UserLog{
		ID:             whitelist.ID,
		EnterTime:      whitelist.EnterTime,
		ExitTime:       whitelist.ExitTime,
		ParkingName:    whitelist.ParkingName,
		ParkingAddress: whitelist.ParkingAddress,
	}
	return response
}

func toUserLogsResSlice(whitelists []entity.UserLog) []entity.UserLog {
	userLogsResSlice := make([]entity.UserLog, 0)
	for _, whitelist := range whitelists {
		userLogsResSlice = append(userLogsResSlice, toUserLogsRes(whitelist))
	}
	return userLogsResSlice
}

type User struct {
	FID        int    `json:"ID"`
	FFirstName string `json:"first_name" validate:"required"`
	FLastName  string `json:"last_name" validate:"required"`
	FCarTag    string `json:"car_tag" validate:"required"`
	FPhone     string `json:"phone" validate:"required"`
	FPassword  string `json:"password" validate:"required"`
}

func (u User) ID() int {
	return u.FID
}

func (u User) FirstName() string {
	return u.FFirstName
}

func (u User) LastName() string {
	return u.FLastName
}

func (u User) CarTag() string {
	return u.FCarTag
}

func (u User) Phone() string {
	return u.FPhone
}

func (u User) Password() string {
	return u.FPassword
}
