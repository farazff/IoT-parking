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

type ParkingRes struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Address   string     `json:"address"`
	Phone     string     `json:"phone"`
	Enabled   bool       `json:"enabled"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func toParkingRes(parking entity.Parking, id int) ParkingRes {
	response := ParkingRes{
		Id:        parking.Id(),
		Name:      parking.Name(),
		Address:   parking.Address(),
		Phone:     parking.Phone(),
		Enabled:   parking.Enabled(),
		CreatedAt: parking.CreatedAt(),
		UpdatedAt: parking.UpdatedAt(),
		DeletedAt: parking.DeletedAt(),
	}
	if id != -1 {
		response.Id = id
	}
	return response
}

func toParkingResSlice(parkings []entity.Parking) []ParkingRes {
	parkingsResSlice := make([]ParkingRes, 0)
	for _, parking := range parkings {
		parkingsResSlice = append(parkingsResSlice, toParkingRes(parking, -1))
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
