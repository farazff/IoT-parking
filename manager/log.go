package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CarEnter(ctx context.Context, carTag string, parkingUUID uuid.UUID) (int, error) {
	userID, err := repository.GetUserIDByCarTag(ctx, carTag)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, fmt.Errorf("error in finding user with given information, %w", err)
	}

	isCarWhiteList, err := repository.IsCarWhitelist(ctx, parkingUUID, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, fmt.Errorf("error in finding withelist with given information, %w", err)
	}
	if !isCarWhiteList {
		return 0, ErrInvalidCarTag
	}

	id, err := repository.CarEnter(ctx, userID, parkingUUID)
	if err != nil {
		lg.Errorf("error during entering car: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func CarExit(ctx context.Context, parkingUUID uuid.UUID, carTag string) error {
	userID, err := repository.GetUserIDByCarTag(ctx, carTag)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding user with given information, %w", err)
	}

	err = repository.CarExit(ctx, parkingUUID, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding parking with given id, %w", err)
	}
	return nil
}

func GetUserLogs(ctx context.Context, phone string) ([]entity.UserLog, error) {
	userID, err := repository.GetUserIDByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	userLogs, err := repository.GetUserLogs(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Whitelists, %w", err)
	}
	return userLogs, nil
}
