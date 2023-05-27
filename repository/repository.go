package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
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
	CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin, uuid uuid.UUID) (int, error)
	GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error)
	GetParkingAdminPasswordByPhone(ctx context.Context, phone string) (string, error)
	GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error)
	UpdateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) error
	DeleteParkingAdmin(ctx context.Context, id int) error
	GetParkingUUID(ctx context.Context, adminUUID uuid.UUID) (uuid.UUID, error)
	GetParkingIdByUuid(ctx context.Context, adminId uuid.UUID) (uuid.UUID, error)
	GetParkingAdminParkingByPhone(ctx context.Context, phone string) (int, error)
}

type ZoneRepository interface {
	CreateZone(ctx context.Context, zone entity.Zone, parkingID int) (int, error)
	GetZones(ctx context.Context, parkingID int) ([]entity.Zone, error)
	GetZone(ctx context.Context, id int, parkingID int) (entity.Zone, error)
	UpdateZone(ctx context.Context, zone entity.Zone, parkingID int) error
	DeleteZone(ctx context.Context, id int, parkingID int) error
	GetCapacitySum(ctx context.Context, id int) (int, error)
	ZoneCarEnter(ctx context.Context, zoneID int) error
	ZoneCarExit(ctx context.Context, zoneID int) error
}

type WhitelistRepository interface {
	CreateWhitelist(ctx context.Context, whitelist entity.Whitelist, parkingUUID uuid.UUID) (int, error)
	GetWhitelists(ctx context.Context, parkingId uuid.UUID) ([]entity.Whitelist, error)
	DeleteWhitelist(ctx context.Context, parkingId uuid.UUID, carTag string) error
	IsCarWhitelist(ct context.Context, parkingId uuid.UUID, carTag string) (bool, error)
}

type LogRepository interface {
	CarEnter(ctx context.Context, log entity.Log) (int, error)
	CarExit(ctx context.Context, pId int, carTag string) error
}
