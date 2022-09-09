package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist) (int, error) {
	id, err := repository.CreateWhitelist(ctx, Whitelist)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating Whitelist: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func GetWhitelists(ctx context.Context) ([]entity.Whitelist, error) {
	Whitelists, err := repository.GetWhitelists(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving Whitelists, %w", err)
	}
	return Whitelists, nil
}

func DeleteWhitelist(ctx context.Context, req entity.WhitelistDeleteReq) error {

	parkingId, err := GetParkingId(ctx, req.AdminCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding parking admin with given id, %w", err)
	}

	err = repository.DeleteWhitelist(ctx, parkingId, req.CarTag)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding Whitelist with given id, %w", err)
	}
	return nil
}
