package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

var ZoneR ZoneRepository

func RegisterZone(p ZoneRepository) error {
	if ZoneR != nil {
		return errors.New("repository: RegisterZone called twice")
	}
	ZoneR = p
	return nil
}

func CreateZone(ctx context.Context, Zone entity.Zone, pUuid uuid.UUID) (int, error) {
	return ZoneR.CreateZone(ctx, Zone, pUuid)
}

func GetZones(ctx context.Context, parkingUUID uuid.UUID) ([]entity.Zone, error) {
	return ZoneR.GetZones(ctx, parkingUUID)
}

func UpdateZone(ctx context.Context, Zone entity.Zone) error {
	return ZoneR.UpdateZone(ctx, Zone)
}

func DeleteZone(ctx context.Context, id int) error {
	return ZoneR.DeleteZone(ctx, id)
}

func GetCapacitySum(ctx context.Context, id int) (int, error) {
	return ZoneR.GetCapacitySum(ctx, id)
}

func GetZone(ctx context.Context, id int) (entity.Zone, error) {
	return ZoneR.GetZone(ctx, id)
}

func ZoneCarEnter(ctx context.Context, ZoneID int) error {
	return ZoneR.ZoneCarEnter(ctx, ZoneID)
}

func ZoneCarExit(ctx context.Context, ZoneID int) error {
	return ZoneR.ZoneCarExit(ctx, ZoneID)
}
