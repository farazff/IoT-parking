package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/google/uuid"
	"github.com/okian/servo/v2/lg"
)

func CreateParking(ctx context.Context, parking entity.Parking) (int, uuid.UUID, error) {
	Puuid := uuid.New()
	id, err := repository.CreateParking(ctx, parking, Puuid.String())
	if err != nil {
		lg.Error("error during creating parking: %v", err)
		return -1, uuid.UUID{}, ErrInternalServer
	}
	return id, Puuid, nil
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

func GetAdminParking(ctx context.Context, phone string) (entity.Parking, int, int, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return nil, 0, 0, err
	}
	parking, capacity, remainedCap, err := repository.GetAdminParking(ctx, parkingID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, 0, 0, ErrNotFound
		}
		return nil, 0, 0, fmt.Errorf("error in retrieving parking, %w", err)
	}
	return parking, capacity, remainedCap, nil
}

func GetParkings(ctx context.Context) ([]entity.Parking, error) {
	parkings, err := repository.GetParkings(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving parkings, %w", err)
	}
	return parkings, nil
}

func GetUserParkings(ctx context.Context, phone string) ([]entity.Parking, error) {
	userID, err := repository.GetUserIDByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	parkings, err := repository.GetUserParkings(ctx, userID)
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
