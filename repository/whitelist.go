package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"

	"github.com/farazff/IoT-parking/entity"
)

var WhitelistR WhitelistRepository

func RegisterWhitelist(p WhitelistRepository) error {
	if WhitelistR != nil {
		return errors.New("repository: RegisterWhitelist called twice")
	}
	WhitelistR = p
	return nil
}

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, parkingID int) (int, error) {
	return WhitelistR.CreateWhitelist(ctx, Whitelist, parkingID)
}

func GetWhitelists(ctx context.Context, parkingID int) ([]entity.Whitelist, error) {
	return WhitelistR.GetWhitelists(ctx, parkingID)
}

func DeleteWhitelist(ctx context.Context, parkingID int, whitelistID int) error {
	return WhitelistR.DeleteWhitelist(ctx, parkingID, whitelistID)
}

func IsCarWhitelist(ctx context.Context, parkingUUID uuid.UUID, carTag string) (bool, error) {
	return WhitelistR.IsCarWhitelist(ctx, parkingUUID, carTag)
}
