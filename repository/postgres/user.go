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
	getUserPasswordByPhone = `SELECT password FROM users WHERE phone = $1`
	getUserIDByPhone       = `SELECT id FROM users WHERE phone = $1`
	getUserIDByCarTag      = `SELECT id FROM users WHERE car_tag = $1`
	createUserQuery        = `INSERT INTO users (first_name, last_name, car_tag, phone, password) VALUES ($1, $2, $3, $4, $5)`
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
	var id int
	err := db.Get(ctx, &id, getUserIDByPhone, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, repository.ErrNotFound
		}
		return -1, err
	}
	return id, nil
}

func (s *service) GetUserIDByCarTag(ctx context.Context, carTag string) (int, error) {
	var id int
	err := db.Get(ctx, &id, getUserIDByCarTag, carTag)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, repository.ErrNotFound
		}
		return -1, err
	}
	return id, nil
}

func (s *service) CreateUser(ctx context.Context, user entity.User) error {
	_, err := db.Exec(ctx, createUserQuery, user.FirstName(), user.LastName(), user.CarTag(), user.Phone(), user.Password())
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return fmt.Errorf("user already exist: %w", repository.ErrDuplicateEntity)
		}
		return err
	}
	return nil
}
