package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
)

type ParkingRepository interface {
	CreateParking(ctx context.Context, parking entity.Parking) (int, error)
	GetParking(ctx context.Context, id int) (entity.Parking, error)
	GetParkings(ctx context.Context) ([]entity.Parking, error)
	UpdateParking(ctx context.Context, parking entity.Parking) error
	DeleteParking(ctx context.Context, id int) error
}

type SystemAdminRepository interface {
	CreateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) (int, error)
	GetSystemAdmin(ctx context.Context, id int) (entity.SystemAdmin, error)
	GetSystemAdmins(ctx context.Context) ([]entity.SystemAdmin, error)
	UpdateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) error
	DeleteSystemAdmin(ctx context.Context, id int) error
}

type ParkingAdminRepository interface {
	CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) (int, error)
	GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error)
	GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error)
	UpdateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) error
	DeleteParkingAdmin(ctx context.Context, id int) error
}
