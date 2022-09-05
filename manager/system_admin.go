package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateSystemAdmin(ctx context.Context, SystemAdmin entity.SystemAdmin) (int, error) {
	id, err := repository.CreateSystemAdmin(ctx, SystemAdmin)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating SystemAdmin: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetSystemAdmin(ctx context.Context, id int) (entity.SystemAdmin, error) {
	SystemAdmin, err := repository.GetSystemAdmin(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error in retrieving SystemAdmin, %w", err)
	}
	return SystemAdmin, nil
}

func GetSystemAdmins(ctx context.Context) ([]entity.SystemAdmin, error) {
	SystemAdmins, err := repository.GetSystemAdmins(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving SystemAdmins, %w", err)
	}
	return SystemAdmins, nil
}

func UpdateSystemAdmin(ctx context.Context, rule entity.SystemAdmin) error {
	err := repository.UpdateSystemAdmin(ctx, rule)
	if err != nil {
		lg.Error("error during updating rule: %v", err)
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return ErrDuplicateEntity
		}
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return ErrInternalServer
	}
	return nil
}

func DeleteSystemAdmin(ctx context.Context, id int) error {
	err := repository.DeleteSystemAdmin(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding SystemAdmin with given id, %w", err)
	}
	return nil
}
