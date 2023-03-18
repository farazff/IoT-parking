package manager

import (
	"context"
	"errors"
	"fmt"

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

func GetZone(ctx context.Context, id int) (entity.Zone, error) {
	Zone, err := repository.GetZone(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error in retrieving Zone, %w", err)
	}
	return Zone, nil
}

func GetZones(ctx context.Context) ([]entity.Zone, error) {
	Zones, err := repository.GetZones(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Zones, %w", err)
	}
	return Zones, nil
}

func UpdateZone(ctx context.Context, zone entity.Zone) error {
	isValid, err := checkAccess(ctx, zone)
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

func DeleteZone(ctx context.Context, id int) error {
	isValid, err := checkAccess(ctx, zone)
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

func checkAccess(ctx context.Context, zone entity.Zone) (bool, error) {
	parkingUUID, err := repository.GetParkingIdByUuid(ctx, zone.AdminUuid())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return false, ErrParkingNotFound
		}
		return false, err
	}

	tempZone, err := repository.GetZone(ctx, zone.Id())
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
