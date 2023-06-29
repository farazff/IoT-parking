package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var userR UserRepository

func RegisterUser(p UserRepository) error {
	if userR != nil {
		return errors.New("repository: RegisterUser called twice")
	}
	userR = p
	return nil
}

func GetUserPasswordByPhone(ctx context.Context, phone string) (string, error) {
	return userR.GetUserPasswordByPhone(ctx, phone)
}

func GetUserIDByPhone(ctx context.Context, phone string) (int, error) {
	return userR.GetUserIDByPhone(ctx, phone)
}

func GetUserIDByCarTag(ctx context.Context, carTag string) (int, error) {
	return userR.GetUserIDByCarTag(ctx, carTag)
}

func CreateUser(ctx context.Context, user entity.User) error {
	return userR.CreateUser(ctx, user)
}
