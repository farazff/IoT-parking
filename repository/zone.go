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

func CreateZone(ctx context.Context, Zone entity.Zone) (int, error) {
	return ZoneR.CreateZone(ctx, Zone)
}

func GetZone(ctx context.Context, id int) (entity.Zone, error) {
	return ZoneR.GetZone(ctx, id)
}

func GetZones(ctx context.Context) ([]entity.Zone, error) {
	return ZoneR.GetZones(ctx)
}

func UpdateZone(ctx context.Context, Zone entity.Zone) error {
	return ZoneR.UpdateZone(ctx, Zone)
}

func DeleteZone(ctx context.Context, id int) error {
	return ZoneR.DeleteZone(ctx, id)
}