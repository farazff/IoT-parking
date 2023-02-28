package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
)

type ParkingRepository interface {
	CreateParking(ctx context.Context, parking entity.Parking, uuid string) (int, error)
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
	GetParkingId(ctx context.Context, adminId int) (int, error)
}

type ZoneRepository interface {
	CreateZone(ctx context.Context, zone entity.Zone) (int, error)
	GetZone(ctx context.Context, id int) (entity.Zone, error)
	GetZones(ctx context.Context) ([]entity.Zone, error)
	UpdateZone(ctx context.Context, zone entity.Zone) error
	DeleteZone(ctx context.Context, id int) error
	GetCapacitySum(ctx context.Context, id int) (int, error)
}

type WhitelistRepository interface {
	CreateWhitelist(ctx context.Context, whitelist entity.Whitelist) (int, error)
	GetWhitelists(ctx context.Context, parkingId int) ([]entity.Whitelist, error)
	DeleteWhitelist(ctx context.Context, parkingId int, carTag string) error
	IsCarWhitelist(ct context.Context, parkingId int, carTag string) (bool, error)
}

type LogRepository interface {
	CarEnter(ctx context.Context, log entity.Log) (int, error)
	CarExit(ctx context.Context, pId int, carTag string) error
}
