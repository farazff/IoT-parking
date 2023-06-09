package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/db"
)

const (
	carEnterQuery = `INSERT INTO logs(car_tag, enter_time, parking_id) VALUES($1, NOW(), $2) RETURNING id`
	carExitQuery  = `UPDATE logs SET (exit_time) = (now()) WHERE parking_id = $1 AND car_tag = $2`
)

func (s *service) CarEnter(ctx context.Context, log entity.Log) (int, error) {
	var id int
	err := db.WQueryRow(ctx, carEnterQuery, log.CarTag(), log.ParkingUUID()).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *service) CarExit(ctx context.Context, pId int, carTag string) error {
	ans, err := db.Exec(ctx, carExitQuery, pId, carTag)
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
