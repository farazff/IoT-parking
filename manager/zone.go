package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateZone(ctx context.Context, zone entity.Zone) (int, error) {
	parkingUUID, err := repository.GetParkingIdByUuid(ctx, zone.AdminUuid())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return -1, ErrParkingNotFound
		}
		return -1, err
	}
	id, err := repository.CreateZone(ctx, zone, parkingUUID)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		if errors.Is(err, repository.ErrParkingForeignKeyConstraint) {
			return id, ErrParkingNotFound
		}
		lg.Error("error during creating Zone: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetZones(ctx context.Context, adminUUID uuid.UUID) ([]entity.Zone, error) {
	parkingUUID, err := repository.GetParkingUUID(ctx, adminUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrParkingNotFound
		}
		return nil, err
	}
	Zones, err := repository.GetZones(ctx, parkingUUID)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Zones, %w", err)
	}
	return Zones, nil
}

func UpdateZone(ctx context.Context, zone entity.Zone) error {
	isValid, err := checkAccess(ctx, zone.AdminUuid(), zone.Id())
	if err != nil {
		return err
	}
	if !isValid {
		return ErrNoAccess
	}

	err = repository.UpdateZone(ctx, zone)
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

func DeleteZone(ctx context.Context, id int, adminUUID uuid.UUID) error {
	isValid, err := checkAccess(ctx, adminUUID, id)
	if err != nil {
		return err
	}
	if !isValid {
		return ErrNoAccess
	}
	err = repository.DeleteZone(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding Zone with given id, %w", err)
	}
	return nil
}

func checkAccess(ctx context.Context, adminUUID uuid.UUID, zoneID int) (bool, error) {
	parkingUUID, err := repository.GetParkingIdByUuid(ctx, adminUUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, ErrParkingNotFound
		}
		return false, err
	}

	tempZone, err := repository.GetZone(ctx, zoneID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, ErrNotFound
		}
		return false, fmt.Errorf("error in retrieving Zone, %w", err)
	}

	if tempZone.PID() != parkingUUID {
		return false, nil
	}
	return true, nil
}

func EnterZone(ctx context.Context, zoneID int) error {
	err := repository.ZoneCarEnter(ctx, zoneID)
	if err != nil {
		return ErrInternalServer
	}
	return nil
}

func ExitZone(ctx context.Context, zoneID int) error {
	err := repository.ZoneCarExit(ctx, zoneID)
	if err != nil {
		return ErrInternalServer
	}
	return nil
}
