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
	createWhitelistQuery = `INSERT INTO Whitelists(id, parking_id, car_tag) VALUES($1, $2, $3) RETURNING id`
	getWhitelistsQuery   = `SELECT id, parking_id, car_tag FROM Whitelists WHERE deleted_at is NULL`
	deleteWhitelistQuery = `DELETE FROM Whitelists where parking_uuid = $1 AND car_tag = $2`
)

func (s *service) CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createWhitelistQuery,
		Whitelist.Id(), Whitelist.PID(), Whitelist.CarTag()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("Whitelist already exist: %w", repository.ErrDuplicateEntity)
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetWhitelists(ctx context.Context) ([]entity.Whitelist, error) {
	var ps []Whitelist
	err := db.Select(ctx, &ps, getWhitelistsQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.Whitelist, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}

func (s *service) DeleteWhitelist(ctx context.Context, req entity.WhitelistDeleteReq) error {
	ans, err := db.Exec(ctx, deleteWhitelistQuery, req.ParkingUuid, req.CarTag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("whitelist doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}
