package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/google/uuid"
)

var logR LogRepository

func RegisterLog(p LogRepository) error {
	if logR != nil {
		return errors.New("repository: RegisterLog called twice")
	}
	logR = p
	return nil
}

func CarEnter(ctx context.Context, userID int, parkingUUID uuid.UUID) (int, error) {
	return logR.CarEnter(ctx, userID, parkingUUID)
}

func CarExit(ctx context.Context, parkingUUID uuid.UUID, userID int) error {
	return logR.CarExit(ctx, parkingUUID, userID)
}

func GetUserLogs(ctx context.Context, userID int, page int, pagination int) ([]entity.UserLog, error) {
	return logR.GetUserLogs(ctx, userID, page, pagination)
}

func GetLogs(ctx context.Context, parkingID int, page int, pagination int) ([]entity.AdminLog, error) {
	return logR.GetLogs(ctx, parkingID, page, pagination)
}
