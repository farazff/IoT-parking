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

func CarEnter(ctx context.Context, log entity.Log, parkingUUID uuid.UUID) (int, error) {
	return logR.CarEnter(ctx, log, parkingUUID)
}

func CarExit(ctx context.Context, parkingUUID uuid.UUID, carTag string) error {
	return logR.CarExit(ctx, parkingUUID, carTag)
}
