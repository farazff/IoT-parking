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
	createParkingQuery = `INSERT INTO parkings(name, address, phone, enabled, created_at, updated_at, uuid) 
							VALUES($1, $2, $3, $4, now(), now(), $5) RETURNING id`
	getParkingsQuery = `SELECT id, name, address, phone, enabled, created_at, updated_at, deleted_at, uuid 
							FROM parkings WHERE deleted_at is NULL`
	getParkingByIdQuery = `SELECT id, name, address, phone, enabled, created_at, updated_at, deleted_at, uuid 
							FROM parkings WHERE deleted_at is NULL AND id = $1`
	updateParkingQuery = `UPDATE parkings SET (name, address, phone, enabled, updated_at) = ($2, $3, $4, $5, now()) 
                			WHERE id = $1 and deleted_at is null`
	deleteParkingQuery = `UPDATE parkings SET deleted_at = now() where id = $1 and deleted_at is null`
)

const uniqueViolation = "23505"

func (s *service) CreateParking(ctx context.Context, parking entity.Parking, uuid string) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createParkingQuery, parking.Name(), parking.Address(), parking.Phone(), parking.Enabled(),
		uuid).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("parking already exist: %w", repository.ErrDuplicateEntity)
		}
		return -1, err
	}
	return id, nil
}

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

func (s *service) UpdateParking(ctx context.Context, parking entity.Parking) error {
	ans, err := db.Exec(ctx, updateParkingQuery,
		parking.Id(), parking.Name(), parking.Address(), parking.Phone(), parking.Enabled())
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("parking already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("parking doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
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
		return fmt.Errorf("parking doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}
