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
	getParkingsQuery = `SELECT id, name, address, phone, enabled, created_at, updated_at, deleted_at 
							FROM parkings WHERE deleted_at is NULL`
	getParkingByIdQuery = `SELECT id, name, address, phone, enabled, created_at, updated_at, deleted_at 
							FROM parkings WHERE deleted_at is NULL AND id = $1`
	deleteParkingQuery = `UPDATE parkings SET deleted_at = now() where id = $1`
)

func (s *service) GetParking(ctx context.Context, id int) (entity.Parking, error) {
	t := Parking{}
	err := db.Get(ctx, &t, getParkingByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *service) GetParkings(ctx context.Context) ([]entity.Parking, error) {
	var ps []Parking
	err := db.Select(ctx, &ps, getParkingsQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.Parking, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}

func (s *service) DeleteParking(ctx context.Context, id int) error {
	ans, err := db.Exec(ctx, deleteParkingQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("rule doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}
