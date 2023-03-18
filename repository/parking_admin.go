package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

var ParkingAdminR ParkingAdminRepository

func RegisterParkingAdmin(p ParkingAdminRepository) error {
	if ParkingAdminR != nil {
		return errors.New("repository: RegisterParkingAdmin called twice")
	}
	ParkingAdminR = p
	return nil
}

func CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin, uuid uuid.UUID) (int, error) {
	return ParkingAdminR.CreateParkingAdmin(ctx, ParkingAdmin, uuid)
}

func GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error) {
	return ParkingAdminR.GetParkingAdmin(ctx, id)
}

func GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error) {
	return ParkingAdminR.GetParkingAdmins(ctx)
}

func UpdateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) error {
	return ParkingAdminR.UpdateParkingAdmin(ctx, ParkingAdmin)
}

func DeleteParkingAdmin(ctx context.Context, id int) error {
	return ParkingAdminR.DeleteParkingAdmin(ctx, id)
}

func GetParkingId(ctx context.Context, adminId int) (uuid.UUID, error) {
	return ParkingAdminR.GetParkingId(ctx, adminId)
}

func GetParkingIdByUuid(ctx context.Context, adminId uuid.UUID) (uuid.UUID, error) {
	return ParkingAdminR.GetParkingIdByUuid(ctx, adminId)
}
