package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/lib/pq"
	"github.com/okian/servo/v2/db"
)

const (
	createZoneQuery = `INSERT INTO Zones(parking_id, capacity, remained_capacity, enabled, created_at, updated_at) 
							VALUES($1, $2, $3, $4, now(), now()) RETURNING id`
	getZonesQuery    = `SELECT id, parking_id, capacity, remained_capacity, enabled FROM Zones WHERE deleted_at is NULL AND parking_id = $1`
	getZoneByIdQuery = `SELECT id, parking_id, capacity, remained_capacity, enabled 
							FROM Zones WHERE deleted_at is NULL AND id = $1 AND parking_id = $2`
	updateZoneQuery = `UPDATE Zones SET (capacity, remained_capacity, enabled, updated_at) = ($2, $3, $4, now()) 
                			WHERE id = $1 and deleted_at is NULL AND parking_id = $5`
	deleteZoneQuery     = `UPDATE Zones SET deleted_at = now() WHERE id = $1 and deleted_at is NULL AND parking_id = $2`
	getCapacitySumQuery = `select sum(capacity) FROM zones WHERE parking_id = $1 and enabled = true`
	carEnterToZoneQuery = `update zones set remained_capacity = remained_capacity - 1
             where id = $1 AND $1 in (select zones.id from zones inner join parkings p on p.id = zones.parking_id where uuid = $2)`
	carExitFromZoneQuery = `update zones set remained_capacity = remained_capacity + 1
             where id = $1 AND $1 in (select zones.id from zones inner join parkings p on p.id = zones.parking_id where uuid = $2)`
)

func (s *service) CreateZone(ctx context.Context, zone entity.Zone, parkingID int) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createZoneQuery,
		parkingID, zone.Capacity(), zone.RemainedCapacity(), zone.Enabled()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == foreignKeyViolation {
			return -1, repository.ErrParkingForeignKeyConstraint
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetZones(ctx context.Context, parkingID int) ([]entity.Zone, error) {
	var ps []Zone
	err := db.Select(ctx, &ps, getZonesQuery, parkingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.Zone, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}

func (s *service) UpdateZone(ctx context.Context, zone entity.Zone, parkingID int) error {
	ans, err := db.Exec(ctx, updateZoneQuery,
		zone.ID(), zone.Capacity(), zone.RemainedCapacity(), zone.Enabled(), parkingID)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("system_admin already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("zone doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteZone(ctx context.Context, id int, parkingID int) error {
	ans, err := db.Exec(ctx, deleteZoneQuery, id, parkingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("zone doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}

func (s *service) GetCapacitySum(ctx context.Context, id int) (int, error) {
	var capacity int
	err := db.Get(ctx, &capacity, getCapacitySumQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("parking not found: %w", repository.ErrNotFound)
		}
		return 0, err
	}
	return capacity, nil
}

func (s *service) GetZone(ctx context.Context, id int, parkingId int) (entity.Zone, error) {
	t := Zone{}
	err := db.Get(ctx, &t, getZoneByIdQuery, id, parkingId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *service) ZoneCarEnter(ctx context.Context, zoneID int, parkingUUID string) error {
	ans, err := db.Exec(ctx, carEnterToZoneQuery, zoneID, parkingUUID)
	if err != nil {
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return repository.ErrNotFound
	}
	return nil
}

func (s *service) ZoneCarExit(ctx context.Context, zoneID int, parkingUUID string) error {
	ans, err := db.Exec(ctx, carExitFromZoneQuery, zoneID, parkingUUID)
	if err != nil {
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return repository.ErrNotFound
	}
	return nil
}
