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
	carEnterQuery = `INSERT INTO logs(car_tag, enter_time, parking_id) VALUES($1, 
                                            NOW(), (SELECT id FROM parkings WHERE uuid = $2)) RETURNING id`
	carExitQuery = `UPDATE logs SET exit_time = NOW() WHERE 
                                          parking_id = (SELECT id FROM parkings WHERE uuid = $1 limit 1) AND car_tag = $2`
)

func (s *service) CarEnter(ctx context.Context, log entity.Log, parkingUUID uuid.UUID) (int, error) {
	var id int
	err := db.WQueryRow(ctx, carEnterQuery, log.CarTag(), parkingUUID).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *service) CarExit(ctx context.Context, parkingUUID uuid.UUID, carTag string) error {
	ans, err := db.Exec(ctx, carExitQuery, parkingUUID, carTag)
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
