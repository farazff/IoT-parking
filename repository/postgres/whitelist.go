package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/okian/servo/v2/db"
)

const (
	createWhitelistQuery = `INSERT INTO whitelists(parking_id, car_tag) VALUES($1, $2) RETURNING id`
	getWhitelistsQuery   = `SELECT id, parking_id, car_tag FROM whitelists WHERE parking_id = $1`
	deleteWhitelistQuery = `DELETE FROM whitelists where parking_id = $1 AND id = $2`
	isCarWhiteListQuery  = `SELECT count(*) from whitelists where parking_id = $1 and car_tag = $2`
)

func (s *service) CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, parkingID int) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createWhitelistQuery, parkingID, Whitelist.CarTag()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("Whitelist already exist: %w", repository.ErrDuplicateEntity)
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetWhitelists(ctx context.Context, parkingID int) ([]entity.Whitelist, error) {
	var ps []Whitelist
	err := db.Select(ctx, &ps, getWhitelistsQuery, parkingID)
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

func (s *service) DeleteWhitelist(ctx context.Context, parkingID int, whiteListID int) error {
	ans, err := db.Exec(ctx, deleteWhitelistQuery, parkingID, whiteListID)
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

func (s *service) IsCarWhitelist(ctx context.Context, parkingId uuid.UUID, carTag string) (bool, error) {
	var count int
	err := db.Get(ctx, &count, isCarWhiteListQuery, parkingId, carTag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("log not found: %w", repository.ErrNotFound)
		}
		return false, err
	}
	if count >= 1 {
		return true, nil
	}
	return false, nil
}
