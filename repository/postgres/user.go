package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/db"
)

const (
	getUserPasswordByPhone = `SELECT password FROM users WHERE deleted_at is NULL AND phone = $1`
	getUserIDByPhone       = `SELECT id FROM users WHERE deleted_at is NULL AND phone = $1`
)

func (s *service) GetUserPasswordByPhone(ctx context.Context, phone string) (string, error) {
	var password string
	err := db.Get(ctx, &password, getUserPasswordByPhone, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", repository.ErrNotFound
		}
		return "", err
	}
	return password, nil
}

func (s *service) GetUserIDByPhone(ctx context.Context, phone string) (int, error) {
	var password int
	err := db.Get(ctx, &password, getUserIDByPhone, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, repository.ErrNotFound
		}
		return -1, err
	}
	return password, nil
}
