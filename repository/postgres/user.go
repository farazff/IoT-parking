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
	getUserPasswordByPhone         = `SELECT password FROM users WHERE phone = $1`
	getUserIDByPhone               = `SELECT id FROM users WHERE phone = $1`
	getUserIDByCarTag              = `SELECT id FROM users WHERE car_tag = $1`
	createUserQuery                = `INSERT INTO users (first_name, last_name, car_tag, phone, password) VALUES ($1, $2, $3, $4, $5)`
	getUserQuery                   = `SELECT first_name, last_name, car_tag, phone FROM users WHERE phone = $1`
	updateUserWithOutPasswordQuery = `UPDATE users SET (first_name, last_name, car_tag) = 
    									($2, $3, $4) WHERE phone = $1`
	updateUserWithPasswordQuery = `UPDATE users SET (first_name, last_name, car_tag, password) = 
    									($2, $3, $4, $5) WHERE phone = $1`
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

func (s *service) GetUser(ctx context.Context, phone string) (entity.User, error) {
	t := User{}
	err := db.Get(ctx, &t, getUserQuery, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return t, nil
}

func (s *service) UpdateUser(ctx context.Context, userUpdater entity.UserUpdater, userPhone string, updatePassword bool) error {
	var ans sql.Result
	var err error
	if !updatePassword {
		ans, err = db.Exec(ctx, updateUserWithOutPasswordQuery, userPhone, userUpdater.FirstName, userUpdater.LastName,
			userUpdater.CarTag)
	} else {
		ans, err = db.Exec(ctx, updateUserWithPasswordQuery, userPhone, userUpdater.FirstName, userUpdater.LastName,
			userUpdater.CarTag, userUpdater.NewPassword)
	}
	if err != nil {
		if err.(*pq.Error).Code == uniqueViolation {
			return repository.ErrDuplicateEntity
		}
		return err
	}
	affected, err := ans.RowsAffected()
	if int(affected) < 1 {
		return fmt.Errorf("user doesn't exist: %w", repository.ErrNotFound)
	}
	if err != nil {
		return err
	}
	return nil
}
