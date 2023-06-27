package repository

import (
	"context"
	"errors"
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
