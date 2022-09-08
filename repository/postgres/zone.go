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
	createZoneQuery = `INSERT INTO Zones(id, parking_id, capacity, remained_capacity, enabled, created_at, updated_at) 
							VALUES($1, $2, $3, $4, $5, now(), now()) RETURNING id`
	getZonesQuery = `SELECT id, parking_id, capacity, remained_capacity, enabled, created_at, updated_at, deleted_at 
							FROM Zones WHERE deleted_at is NULL`
	getZoneByIdQuery = `SELECT id, parking_id, capacity, remained_capacity, enabled, created_at, updated_at, deleted_at 
							FROM Zones WHERE deleted_at is NULL AND id = $1`
	updateZoneQuery = `UPDATE Zones SET (parking_id, capacity, remained_capacity, enabled, updated_at) = ($2, $3, $4, $5, now()) 
                			WHERE id = $1`
	deleteZoneQuery     = `UPDATE Zones SET deleted_at = now() WHERE id = $1`
	getCapacitySumQuery = `select sum(capacity) FROM zones WHERE parking_id = $1 and enabled = true`
)

func (s *service) CreateZone(ctx context.Context, Zone entity.Zone) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createZoneQuery,
		Zone.Id(), Zone.PID(), Zone.Capacity(), Zone.RemainedCapacity(), Zone.Enabled()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("Zone already exist: %w", repository.ErrDuplicateEntity)
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetZone(ctx context.Context, id int) (entity.Zone, error) {
	t := Zone{}
	err := db.Get(ctx, &t, getZoneByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *service) GetZones(ctx context.Context) ([]entity.Zone, error) {
	var ps []Zone
	err := db.Select(ctx, &ps, getZonesQuery)
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

func (s *service) UpdateZone(ctx context.Context, Zone entity.Zone) error {
	ans, err := db.Exec(ctx, updateZoneQuery,
		Zone.Id(), Zone.PID(), Zone.Capacity(), Zone.RemainedCapacity(), Zone.Enabled())
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("system_admin already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("system_admin doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteZone(ctx context.Context, id int) error {
	ans, err := db.Exec(ctx, deleteZoneQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("system_admin doesn't exist: %w", repository.ErrNotFound)
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
