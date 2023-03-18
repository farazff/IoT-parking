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
	createParkingAdminQuery = `INSERT INTO parking_admins(first_name, last_name, phone, parking_id, enabled, created_at, updated_at, uuid) 
							VALUES($1, $2, $3, $4, $5, now(), now(), $6) RETURNING id`
	getParkingAdminsQuery = `SELECT id, first_name, last_name, phone, parking_id, enabled 
							FROM parking_admins WHERE deleted_at is NULL`
	getParkingAdminByIdQuery = `SELECT id, first_name, last_name, phone, parking_id, enabled 
							FROM parking_admins WHERE deleted_at is NULL AND id = $1`
	updateParkingAdminQuery = `UPDATE parking_admins SET (first_name, last_name, phone, parking_id, enabled, updated_at)= ($2, $3, $4, $5, $6, now()) 
                			WHERE id = $1 and deleted_at is null`
	deleteParkingAdminQuery = `UPDATE parking_admins SET deleted_at = now() where id = $1 and deleted_at is null`
	getParkingIdQuery       = `select parking_id from parking_admins where id = $1`
	getParkingIdQueryByUuid = `select parking_id from parking_admins where uuid = $1`
)

func (s *service) CreateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin, uuid uuid.UUID) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createParkingAdminQuery, ParkingAdmin.FirstName(), ParkingAdmin.LastName(),
		ParkingAdmin.Phone(), ParkingAdmin.PID(), ParkingAdmin.Enabled(), uuid).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("ParkingAdmin already exist: %w", repository.ErrDuplicateEntity)
		}
		if err.(*pq.Error).Code == foreignKeyViolation {
			return -1, repository.ErrParkingForeignKeyConstraint
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetParkingAdmin(ctx context.Context, id int) (entity.ParkingAdmin, error) {
	t := ParkingAdmin{}
	err := db.Get(ctx, &t, getParkingAdminByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *service) GetParkingAdmins(ctx context.Context) ([]entity.ParkingAdmin, error) {
	var ps []ParkingAdmin
	err := db.Select(ctx, &ps, getParkingAdminsQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.ParkingAdmin, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}

func (s *service) UpdateParkingAdmin(ctx context.Context, ParkingAdmin entity.ParkingAdmin) error {
	ans, err := db.Exec(ctx, updateParkingAdminQuery,
		ParkingAdmin.Id(), ParkingAdmin.FirstName(), ParkingAdmin.LastName(), ParkingAdmin.Phone(), ParkingAdmin.PID(), ParkingAdmin.Enabled())
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("parking_admin already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("parking_admin doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteParkingAdmin(ctx context.Context, id int) error {
	ans, err := db.Exec(ctx, deleteParkingAdminQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("parking_admin doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}

func (s *service) GetParkingId(ctx context.Context, AdminId int) (uuid.UUID, error) {
	var parkingId uuid.UUID
	err := db.Get(ctx, &parkingId, getParkingIdQuery, AdminId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, fmt.Errorf("parking admin not found: %w", repository.ErrNotFound)
		}
		return uuid.UUID{}, err
	}
	return parkingId, nil
}

func (s *service) GetParkingIdByUuid(ctx context.Context, AdminId uuid.UUID) (uuid.UUID, error) {
	var parkingId uuid.UUID
	err := db.Get(ctx, &parkingId, getParkingIdQueryByUuid, AdminId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.UUID{}, fmt.Errorf("parking admin not found: %w", repository.ErrNotFound)
		}
		return uuid.UUID{}, err
	}
	return parkingId, nil
}
