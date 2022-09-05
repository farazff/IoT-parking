package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateParking(ctx context.Context, parking entity.Parking) (int, error) {
	id, err := repository.CreateParking(ctx, parking)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating parking: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetParking(ctx context.Context, id int) (entity.Parking, int, error) {
	parking, err := repository.GetParking(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, 0, ErrNotFound
		}
		return nil, 0, fmt.Errorf("error in retrieving parking, %w", err)
	}
	parkingCapacity, err := repository.GetCapacitySum(ctx, id)
	return parking, parkingCapacity, nil
}

func GetParkings(ctx context.Context) ([]entity.Parking, error) {
	parkings, err := repository.GetParkings(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving parkings, %w", err)
	}
	return parkings, nil
}

func UpdateParking(ctx context.Context, rule entity.Parking) error {
	err := repository.UpdateParking(ctx, rule)
	if err != nil {
		lg.Error("error during updating rule: %v", err)
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return ErrDuplicateEntity
		}
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternalServer
	}
	return nil
}

func DeleteParking(ctx context.Context, id int) error {
	err := repository.DeleteParking(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding parking with given id, %w", err)
	}
	return nil
}
