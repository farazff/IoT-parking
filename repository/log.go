package repository

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
)

var logR LogRepository

func RegisterLog(p LogRepository) error {
	if logR != nil {
		return errors.New("repository: RegisterLog called twice")
	}
	logR = p
	return nil
}

func CarEnter(ctx context.Context, log entity.Log) (int, error) {
	return logR.CarEnter(ctx, log)
}

func CarExit(ctx context.Context, pId int, carTag string) error {
	return logR.CarExit(ctx, pId, carTag)
}
