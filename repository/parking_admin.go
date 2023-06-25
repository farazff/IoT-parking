package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var ParkingAdminR ParkingAdminRepository

func RegisterParkingAdmin(p ParkingAdminRepository) error {
	if ParkingAdminR != nil {
		return errors.New("repository: RegisterParkingAdmin called twice")
	}
	ParkingAdminR = p
	return nil
}

func CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) (int, error) {
	return ParkingAdminR.CreateParkingAdmin(ctx, ParkingAdmin)
}

func GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error) {
	return ParkingAdminR.GetParkingAdmin(ctx, id)
}

func GetParkingAdminPasswordByPhone(ctx context.Context, phone string) (string, error) {
	return ParkingAdminR.GetParkingAdminPasswordByPhone(ctx, phone)
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

func GetParkingAdminParkingByPhone(ctx context.Context, phone string) (int, error) {
	return ParkingAdminR.GetParkingAdminParkingByPhone(ctx, phone)
}
