package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateUser(ctx context.Context, user entity.User) error {
	err := repository.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return ErrDuplicateEntity
		}
		lg.Error("error during creating user: %v", err)
		return ErrInternalServer
	}
	return nil
}

func GetUser(ctx context.Context, userPhone string) (entity.User, error) {
	user, err := repository.GetUser(ctx, userPhone)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error in retrieving User, %w", err)
	}
	return user, nil
}

func UpdateUser(ctx context.Context, userUpdater entity.UserUpdater, userPhone string) error {
	updatePassword := false
	if len(userUpdater.NewPassword) != 0 && len(userUpdater.OldPassword) != 0 {
		updatePassword = true
		password, err := repository.GetUserPasswordByPhone(ctx, userPhone)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				return ErrNotFound
			}
			return fmt.Errorf("error in retrieving user, %w", err)
		}
		if password != userUpdater.OldPassword {
			return ErrUnauthorized
		}
	}

	err := repository.UpdateUser(ctx, userUpdater, userPhone, updatePassword)
	if err != nil {
		lg.Error("error during updating user: %v", err)
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return ErrDuplicateEntity
		}
		if errors.Is(err, repository.ErrParkingForeignKeyConstraint) {
			return ErrParkingNotFound
		}
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternalServer
	}
	return nil
}
