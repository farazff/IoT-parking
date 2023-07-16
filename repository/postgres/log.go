package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/google/uuid"
	"github.com/okian/servo/v2/db"
)

const (
	carEnterQuery = `INSERT INTO logs(user_id, enter_time, parking_id) VALUES($1, 
                                            NOW(), (SELECT id FROM parkings WHERE uuid = $2)) RETURNING id`

	carExitQuery = `UPDATE logs SET exit_time = NOW() WHERE 
                                          parking_id = (SELECT id FROM parkings WHERE uuid = $1 limit 1) AND user_id = $2`

	GetUserLogsQuery = `SELECT l.id AS id, l.enter_time AS enter_time, l.exit_time AS exit_time, p.name AS parking_name,
					p.address AS parking_address FROM logs AS l JOIN parkings AS p 
					    on l.parking_id = p.id WHERE l.user_id = $1 ORDER BY id DESC`

	GetLogsQuery = `SELECT l.id AS id, l.enter_time AS enter_time, l.exit_time AS exit_time, 
       u.first_name as first_name, u.last_name as last_name, u.car_tag as car_tag, u.phone as phone FROM logs AS l JOIN users AS u 
					    on l.user_id = u.id WHERE l.parking_id = $1 ORDER BY id DESC`
)

func (s *service) CarEnter(ctx context.Context, userID int, parkingUUID uuid.UUID) (int, error) {
	var id int
	err := db.WQueryRow(ctx, carEnterQuery, userID, parkingUUID).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *service) CarExit(ctx context.Context, parkingUUID uuid.UUID, userID int) error {
	ans, err := db.Exec(ctx, carExitQuery, parkingUUID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("parking doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}

func (s *service) GetUserLogs(ctx context.Context, userID int, page int, pagination int) ([]entity.UserLog, error) {
	var wls []entity.UserLog
	err := db.Select(ctx, &wls, GetUserLogsQuery, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.UserLog, 0)
	for i := range wls {
		res = append(res, wls[i])
	}
	return res, nil
}

func (s *service) GetLogs(ctx context.Context, parkingID int, page int, pagination int) ([]entity.AdminLog, error) {
	var wls []entity.AdminLog
	err := db.Select(ctx, &wls, GetLogsQuery, parkingID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.AdminLog, 0)
	for i := range wls {
		res = append(res, wls[i])
	}
	return res, nil
}
