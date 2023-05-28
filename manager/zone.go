package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateZone(ctx context.Context, zone entity.Zone, phone string) (int, int, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return -1, -1, err
	}
	id, err := repository.CreateZone(ctx, zone, parkingID)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, -1, ErrDuplicateEntity
		}
		if errors.Is(err, repository.ErrParkingForeignKeyConstraint) {
			return id, -1, ErrParkingNotFound
		}
		lg.Error("error during creating Zone: %v", err)
		return id, -1, ErrInternalServer
	}
	return id, parkingID, nil
}

func GetZones(ctx context.Context, phone string) ([]entity.Zone, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	Zones, err := repository.GetZones(ctx, parkingID)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Zones, %w", err)
	}
	return Zones, nil
}

func GetZone(ctx context.Context, zoneID int, phone string) (entity.Zone, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}
	Zones, err := repository.GetZone(ctx, zoneID, parkingID)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Zones, %w", err)
	}
	return Zones, nil
}

func UpdateZone(ctx context.Context, zone entity.Zone, phone string) error {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return err
	}

	err = repository.UpdateZone(ctx, zone, parkingID)
	if err != nil {
		lg.Error("error during updating zone: %v", err)
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

func DeleteZone(ctx context.Context, id int, phone string) error {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return err
	}
	err = repository.DeleteZone(ctx, id, parkingID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding Zone with given id, %w", err)
	}
	return nil
}

func EnterZone(ctx context.Context, zoneID int, parkingUUID string) error {

	err := repository.ZoneCarEnter(ctx, zoneID, parkingUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternalServer
	}
	return nil
}

func ExitZone(ctx context.Context, zoneID int, parkingUUID string) error {
	err := repository.ZoneCarExit(ctx, zoneID, parkingUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternalServer
	}
	return nil
}
