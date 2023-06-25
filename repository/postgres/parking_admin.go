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
	createParkingAdminQuery = `INSERT INTO parking_admins(first_name, last_name, phone, enabled, created_at, updated_at, password, parking_id) 
							VALUES($1, $2, $3, $4, now(), now(), $5, $6) RETURNING id`
	getParkingAdminsQuery = `SELECT id, first_name, last_name, phone, enabled, password, parking_id 
							FROM parking_admins WHERE deleted_at is NULL`
	getParkingAdminByIdQuery = `SELECT id, first_name, last_name, phone, enabled, password, parking_id 
							FROM parking_admins WHERE deleted_at is NULL AND id = $1`
	updateParkingAdminQuery = `UPDATE parking_admins SET (first_name, last_name, phone, enabled, updated_at, password, parking_id) = 
    									($2, $3, $4, $5, now(), $6, $7) WHERE id = $1 and deleted_at is null`
	deleteParkingAdminQuery        = `UPDATE parking_admins SET deleted_at = now() where id = $1 and deleted_at is null`
	getParkingAdminPasswordByPhone = `SELECT password FROM parking_admins WHERE deleted_at is NULL AND phone = $1`
	getParkingAdminParkingByPhone  = `SELECT parking_id FROM parking_admins WHERE deleted_at is NULL AND phone = $1`
)

func (s *service) CreateParkingAdmin(ctx context.Context, parkingAdmin entity.ParkingAdmin) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createParkingAdminQuery, parkingAdmin.FirstName(), parkingAdmin.LastName(),
		parkingAdmin.Phone(), parkingAdmin.Enabled(), parkingAdmin.Password(), parkingAdmin.ParkingID()).Scan(&id)
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

func (s *service) GetParkingAdminPasswordByPhone(ctx context.Context, phone string) (string, error) {
	var password string
	err := db.Get(ctx, &password, getParkingAdminPasswordByPhone, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.ErrNotFound
		}
		return "", err
	}

	return password, nil
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
		ParkingAdmin.ID(), ParkingAdmin.FirstName(), ParkingAdmin.LastName(), ParkingAdmin.Phone(), ParkingAdmin.ParkingID(), ParkingAdmin.Enabled())
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

func (s *service) GetParkingAdminParkingByPhone(ctx context.Context, phone string) (int, error) {
	var parkingID int
	err := db.Get(ctx, &parkingID, getParkingAdminParkingByPhone, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, repository.ErrNotFound
		}
		return -1, err
	}
	return parkingID, nil
}
