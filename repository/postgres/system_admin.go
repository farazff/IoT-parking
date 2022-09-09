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
	createSystemAdminQuery = `INSERT INTO SystemAdmins(first_name, last_name, phone, enabled, created_at, updated_at) 
							VALUES($1, $2, $3, $4, now(), now()) RETURNING id`
	getSystemAdminsQuery = `SELECT id, first_name, last_name, phone, enabled, created_at, updated_at, deleted_at 
							FROM SystemAdmins WHERE deleted_at is NULL`
	getSystemAdminByIdQuery = `SELECT id, first_name, last_name, phone, enabled, created_at, updated_at, deleted_at 
							FROM SystemAdmins WHERE deleted_at is NULL AND id = $1`
	updateSystemAdminQuery = `UPDATE SystemAdmins SET (first_name, last_name, phone, enabled, updated_at) = ($2, $3, $4, $5, now()) 
                			WHERE id = $1`
	deleteSystemAdminQuery = `UPDATE SystemAdmins SET deleted_at = now() where id = $1`
)

func (s *service) CreateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) (int, error) {
	var id int
	err := db.WQueryRow(ctx, createSystemAdminQuery, SystemAdmin.FirstName(), SystemAdmin.LastName(),
		SystemAdmin.Phone(), SystemAdmin.Enabled()).Scan(&id)
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return -1, fmt.Errorf("SystemAdmin already exist: %w", repository.ErrDuplicateEntity)
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetSystemAdmin(ctx context.Context, id int) (entity.SystemAdmin, error) {
	t := SystemAdmin{}
	err := db.Get(ctx, &t, getSystemAdminByIdQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return t, nil
}

func (s *service) GetSystemAdmins(ctx context.Context) ([]entity.SystemAdmin, error) {
	var ps []SystemAdmin
	err := db.Select(ctx, &ps, getSystemAdminsQuery)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	res := make([]entity.SystemAdmin, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}

func (s *service) UpdateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) error {
	ans, err := db.Exec(ctx, updateSystemAdminQuery,
		SystemAdmin.Id(), SystemAdmin.FirstName(), SystemAdmin.LastName(), SystemAdmin.Phone(), SystemAdmin.Enabled())
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("system_admin already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("system_admin doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteSystemAdmin(ctx context.Context, id int) error {
	ans, err := db.Exec(ctx, deleteSystemAdminQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("system_admin doesn't exist: %w", repository.ErrNotFound)
	}
	return nil
}
