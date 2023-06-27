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

func ApproveWhitelist(ctx context.Context, whitelistID int, parkingID int) error {
	return WhitelistR.ApproveWhitelist(ctx, whitelistID, parkingID)
}

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, userID int) (int, error) {
	return WhitelistR.CreateWhitelist(ctx, Whitelist, userID)
}

func GetWhitelists(ctx context.Context, parkingID int, approved bool) ([]entity.WhitelistOfficeData, error) {
	return WhitelistR.GetWhitelists(ctx, parkingID, approved)
}

func DeleteWhitelist(ctx context.Context, parkingID int, whitelistID int) error {
	return WhitelistR.DeleteWhitelist(ctx, parkingID, whitelistID)
}

func IsCarWhitelist(ctx context.Context, parkingUUID uuid.UUID, carTag string) (bool, error) {
	return WhitelistR.IsCarWhitelist(ctx, parkingUUID, carTag)
}

func GetUserWhitelists(ctx context.Context, userID int) ([]entity.WhitelistUserData, error) {
	return WhitelistR.GetUserWhitelists(ctx, userID)
}
