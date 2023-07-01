package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/lib/pq"
	"github.com/okian/servo/v2/db"
)

const (
	approveWhiteListQuery = `UPDATE whitelists SET approved = TRUE WHERE id = $1 AND parking_id = $2 AND approved = false`
	createWhitelistQuery  = `INSERT INTO whitelists(user_id, parking_id) VALUES($1, $2) RETURNING id`
	getWhitelistsQuery    = `SELECT w.id as id, u.first_name as first_name, u.last_name as last_name, 
       									u.car_tag as car_tag FROM whitelists 
    								as w join users as u on w.user_id = u.id WHERE parking_id = $1 AND approved = $2`
	deleteWhitelistQuery = `DELETE FROM whitelists where parking_id = $1 AND id = $2`
	isCarWhiteListQuery  = `SELECT count(*) from whitelists where 
                            	parking_id = (SELECT id from parkings where uuid = $1) and user_id = $2`
	getUserWhitelistsQuery = `SELECT w.id as id, p.name as parking_name, p.address as parking_address, 
       								w.approved as approved FROM whitelists 
    										as w join parkings as p on w.parking_id = p.id WHERE w.user_id = $1`
)

func (s *service) ApproveWhitelist(ctx context.Context, whiteListID int, parkingID int) error {
	ans, err := db.Exec(ctx, approveWhiteListQuery, whiteListID, parkingID)
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

func (s *service) CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, userID int) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createWhitelistQuery, userID, Whitelist.ParkingID()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, repository.ErrDuplicateEntity
		}
		if err.(*pq.Error).Code == foreignKeyViolation {
			return -1, repository.ErrParkingForeignKeyConstraint
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetWhitelists(ctx context.Context, parkingID int, approved bool) ([]entity.WhitelistOfficeData, error) {
	var ps []entity.WhitelistOfficeData
	err := db.Select(ctx, &ps, getWhitelistsQuery, parkingID, approved)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.WhitelistOfficeData, 0)
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

func (s *service) IsCarWhitelist(ctx context.Context, parkingUUID uuid.UUID, userID int) (bool, error) {
	var count int
	err := db.Get(ctx, &count, isCarWhiteListQuery, parkingUUID, userID)
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

func (s *service) GetUserWhitelists(ctx context.Context, userID int) ([]entity.WhitelistUserData, error) {
	var ps []entity.WhitelistUserData
	err := db.Select(ctx, &ps, getUserWhitelistsQuery, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.WhitelistUserData, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}
