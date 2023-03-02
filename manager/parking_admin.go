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

func CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) (int, error) {
	id, err := repository.CreateParkingAdmin(ctx, ParkingAdmin)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating ParkingAdmin: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error) {
	ParkingAdmin, err := repository.GetParkingAdmin(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error in retrieving ParkingAdmin, %w", err)
	}
	return ParkingAdmin, nil
}

func GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error) {
	ParkingAdmins, err := repository.GetParkingAdmins(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving ParkingAdmins, %w", err)
	}
	return ParkingAdmins, nil
}

func UpdateParkingAdmin(ctx context.Context, rule entity.ParkingAdmin) error {
	err := repository.UpdateParkingAdmin(ctx, rule)
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

func DeleteParkingAdmin(ctx context.Context, id int) error {
	err := repository.DeleteParkingAdmin(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding ParkingAdmin with given id, %w", err)
	}
	return nil
}

func GetParkingId(ctx context.Context, adminId int) (uuid.UUID, error) {
	parkingId, err := repository.GetParkingId(ctx, adminId)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return uuid.UUID{}, ErrNotFound
		}
		return uuid.UUID{}, fmt.Errorf("error in finding ParkingAdmin with given id, %w", err)
	}
	return parkingId, nil
}
