package manager

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/google/uuid"
	"github.com/spf13/viper"

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

func GetUserLogs(ctx context.Context, phone string, page int) ([]entity.UserLog, int, error) {
	userID, err := repository.GetUserIDByPhone(ctx, phone)
	if err != nil {
		return nil, 0, err
	}

	pagination := viper.GetInt("logs_pagination")
	userLogs, err := repository.GetUserLogs(ctx, userID, page, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("error in retrieving user logs, %w", err)
	}
	offset := pagination * (page - 1)
	lenAll := len(userLogs)
	if len(userLogs) > offset {
		userLogs = userLogs[offset:int(math.Min(float64(len(userLogs)), float64(offset+pagination)))]
	} else {
		userLogs = make([]entity.UserLog, 0)
	}

	pageCount := 0
	if lenAll%pagination == 0 {
		pageCount = lenAll / pagination
	} else {
		pageCount = (lenAll / pagination) + 1
	}

	return userLogs, pageCount, nil
}

func GetLogs(ctx context.Context, phone string, page int) ([]entity.AdminLog, int, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return nil, 0, err
	}

	pagination := viper.GetInt("logs_pagination")
	adminLogs, err := repository.GetLogs(ctx, parkingID, page, pagination)
	if err != nil {
		return nil, 0, fmt.Errorf("error in retrieving logs, %w", err)
	}
	offset := pagination * (page - 1)
	lenAll := len(adminLogs)
	if len(adminLogs) > offset {
		adminLogs = adminLogs[offset:int(math.Min(float64(len(adminLogs)), float64(offset+pagination)))]
	} else {
		adminLogs = make([]entity.AdminLog, 0)
	}

	pageCount := 0
	if lenAll%pagination == 0 {
		pageCount = lenAll / pagination
	} else {
		pageCount = (lenAll / pagination) + 1
	}

	return adminLogs, pageCount, nil
}
