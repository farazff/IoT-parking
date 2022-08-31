package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
)

func GetParking(ctx context.Context, id int) (entity.Parking, error) {
	parking, err := repository.GetParking(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrParkingNotFound
		}
		return nil, fmt.Errorf("error in retrieving parking, %w", err)
	}
	return parking, nil
}

func GetParkings(ctx context.Context) ([]entity.Parking, error) {
	parkings, err := repository.GetParkings(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving parkings, %w", err)
	}
	return parkings, nil
}

func DeleteParking(ctx context.Context, id int) error {
	err := repository.DeleteParking(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrParkingNotFound
		}
		return fmt.Errorf("error in finding parking with given id, %w", err)
	}
	return nil
}
