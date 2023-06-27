package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/db"
)

const (
	getUserPasswordByPhone = `SELECT password FROM system_admins WHERE deleted_at is NULL AND phone = $1`
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
