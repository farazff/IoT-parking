package rest

import (
	"github.com/farazff/IoT-parking/entity"
	"time"
)

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

func toParkingRes(parking entity.Parking) ParkingRes {
	return ParkingRes{
		Id:        parking.Id(),
		Name:      parking.Name(),
		Address:   parking.Address(),
		Phone:     parking.Phone(),
		Enabled:   parking.Enabled(),
		CreatedAt: parking.CreatedAt(),
		UpdatedAt: parking.UpdatedAt(),
		DeletedAt: parking.DeletedAt(),
	}
}

func toParkingResSlice(parkings []entity.Parking) []ParkingRes {
	parkingsResSlice := make([]ParkingRes, 0)
	for _, parking := range parkings {
		parkingsResSlice = append(parkingsResSlice, toParkingRes(parking))
	}
	return parkingsResSlice
}
