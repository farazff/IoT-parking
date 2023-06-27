package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func ApproveWhitelist(ctx context.Context, whiteListID int, phone string) error {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return err
	}

	err = repository.ApproveWhitelist(ctx, whiteListID, parkingID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding Whitelist with given id, %w", err)
	}
	return nil
}

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, phone string) (int, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return -1, err
	}

	id, err := repository.CreateWhitelist(ctx, Whitelist, parkingID)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating Whitelist: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetWhitelists(ctx context.Context, phone string, approved bool) ([]entity.WhitelistOfficeData, error) {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return nil, err
	}

	whitelistsOfficeData, err := repository.GetWhitelists(ctx, parkingID, approved)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Whitelists, %w", err)
	}
	return whitelistsOfficeData, nil
}

func DeleteWhitelist(ctx context.Context, whiteListID int, phone string) error {
	parkingID, err := repository.GetParkingAdminParkingByPhone(ctx, phone)
	if err != nil {
		return err
	}

	err = repository.DeleteWhitelist(ctx, parkingID, whiteListID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding Whitelist with given id, %w", err)
	}
	return nil
}
