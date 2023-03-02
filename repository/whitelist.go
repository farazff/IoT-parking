package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

var WhitelistR WhitelistRepository

func RegisterWhitelist(p WhitelistRepository) error {
	if WhitelistR != nil {
		return errors.New("repository: RegisterWhitelist called twice")
	}
	WhitelistR = p
	return nil
}

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist) (int, error) {
	return WhitelistR.CreateWhitelist(ctx, Whitelist)
}

func GetWhitelists(ctx context.Context, parkingId uuid.UUID) ([]entity.Whitelist, error) {
	return WhitelistR.GetWhitelists(ctx, parkingId)
}

func DeleteWhitelist(ctx context.Context, parkingId uuid.UUID, carTag string) error {
	return WhitelistR.DeleteWhitelist(ctx, parkingId, carTag)
}

func IsCarWhitelist(ctx context.Context, parkingId uuid.UUID, carTag string) (bool, error) {
	return WhitelistR.IsCarWhitelist(ctx, parkingId, carTag)
}
