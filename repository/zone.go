package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var ZoneR ZoneRepository

func RegisterZone(p ZoneRepository) error {
	if ZoneR != nil {
		return errors.New("repository: RegisterZone called twice")
	}
	ZoneR = p
	return nil
}

func CreateZone(ctx context.Context, zone entity.Zone, parkingID int) (int, error) {
	return ZoneR.CreateZone(ctx, zone, parkingID)
}

func GetZones(ctx context.Context, parkingID int) ([]entity.Zone, error) {
	return ZoneR.GetZones(ctx, parkingID)
}

func UpdateZone(ctx context.Context, Zone entity.Zone, parkingID int) error {
	return ZoneR.UpdateZone(ctx, Zone, parkingID)
}

func DeleteZone(ctx context.Context, id int, parkingID int) error {
	return ZoneR.DeleteZone(ctx, id, parkingID)
}

func GetCapacitySum(ctx context.Context, id int) (int, error) {
	return ZoneR.GetCapacitySum(ctx, id)
}

func GetZone(ctx context.Context, id int, parkingID int) (entity.Zone, error) {
	return ZoneR.GetZone(ctx, id, parkingID)
}

func ZoneCarEnter(ctx context.Context, ZoneID int) error {
	return ZoneR.ZoneCarEnter(ctx, ZoneID)
}

func ZoneCarExit(ctx context.Context, ZoneID int) error {
	return ZoneR.ZoneCarExit(ctx, ZoneID)
}
