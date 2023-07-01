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
	GetUserParkings(ctx context.Context, userID int) ([]entity.Parking, error)
	UpdateParking(ctx context.Context, parking entity.Parking) error
	DeleteParking(ctx context.Context, id int) error
}

type SystemAdminRepository interface {
	CreateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) (int, error)
	GetSystemAdmin(ctx context.Context, id int) (entity.SystemAdmin, error)
	GetSystemAdmins(ctx context.Context) ([]entity.SystemAdmin, error)
	UpdateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) error
	DeleteSystemAdmin(ctx context.Context, id int) error
	GetSystemAdminPasswordByPhone(ctx context.Context, phone string) (string, error)
}

type ParkingAdminRepository interface {
	CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) (int, error)
	GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error)
	GetParkingAdminPasswordByPhone(ctx context.Context, phone string) (string, error)
	GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error)
	UpdateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) error
	DeleteParkingAdmin(ctx context.Context, id int) error
	GetParkingAdminParkingByPhone(ctx context.Context, phone string) (int, error)
}

type ZoneRepository interface {
	CreateZone(ctx context.Context, zone entity.Zone, parkingID int) (int, error)
	GetZones(ctx context.Context, parkingID int) ([]entity.Zone, error)
	GetZone(ctx context.Context, id int, parkingID int) (entity.Zone, error)
	UpdateZone(ctx context.Context, zone entity.Zone, parkingID int) error
	DeleteZone(ctx context.Context, id int, parkingID int) error
	GetCapacitySum(ctx context.Context, id int) (int, error)
	ZoneCarEnter(ctx context.Context, zoneID int, parkingUUID string) error
	ZoneCarExit(ctx context.Context, zoneID int, parkingUUID string) error
}

type WhitelistRepository interface {
	ApproveWhitelist(ctx context.Context, whiteListID int, parkingID int) error
	CreateWhitelist(ctx context.Context, whitelist entity.Whitelist, userID int) (int, error)
	GetWhitelists(ctx context.Context, parkingID int, approved bool) ([]entity.WhitelistOfficeData, error)
	DeleteWhitelist(ctx context.Context, parkingID int, whiteListID int) error
	IsCarWhitelist(ct context.Context, parkingUUID uuid.UUID, userID int) (bool, error)
	GetUserWhitelists(ctx context.Context, userID int) ([]entity.WhitelistUserData, error)
}

type LogRepository interface {
	CarEnter(ctx context.Context, userID int, parkingUUID uuid.UUID) (int, error)
	CarExit(ctx context.Context, parkingUUID uuid.UUID, userID int) error
	GetUserLogs(ctx context.Context, userID int) ([]entity.UserLog, error)
}

type UserRepository interface {
	GetUserPasswordByPhone(ctx context.Context, phone string) (string, error)
	GetUserIDByPhone(ctx context.Context, phone string) (int, error)
	GetUserIDByCarTag(ctx context.Context, carTag string) (int, error)
	CreateUser(ctx context.Context, user entity.User) error
}
