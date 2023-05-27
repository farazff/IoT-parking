package manager

import (
	"context"
	"errors"
	"fmt"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CarEnter(ctx context.Context, log entity.Log) (int, error) {

	isCarWhiteList, err := repository.IsCarWhitelist(ctx, log.PID(), log.CarTag())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, fmt.Errorf("error in finding withelist with given information, %w", err)
	}
	if !isCarWhiteList {
		return 0, ErrInvalidCarTag
	}

	id, err := repository.CarEnter(ctx, log)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Errorf("error during creating parking: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

func CarExit(ctx context.Context, pId int, carTag string) error {
	err := repository.CarExit(ctx, pId, carTag)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("error in finding parking with given id, %w", err)
	}
	return nil
}
