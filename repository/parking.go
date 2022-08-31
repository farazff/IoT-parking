package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var parkingR ParkingRepository

func RegisterParking(p ParkingRepository) error {
	if parkingR != nil {
		return errors.New("repository: RegisterParking called twice")
	}
	parkingR = p
	return nil
}

func GetParking(ctx context.Context, id int) (entity.Parking, error) {
	return parkingR.GetParking(ctx, id)
}

func GetParkings(ctx context.Context) ([]entity.Parking, error) {
	return parkingR.GetParkings(ctx)
}

func DeleteParking(ctx context.Context, id int) error {
	return parkingR.DeleteParking(ctx, id)
}
