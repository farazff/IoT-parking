package manager

import (
	"context"
	"errors"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/okian/servo/v2/lg"
)

func CreateUser(ctx context.Context, user entity.User) error {
	err := repository.CreateUser(ctx, user)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return ErrDuplicateEntity
		}
		lg.Error("error during creating user: %v", err)
		return ErrInternalServer
	}
	return nil
}
