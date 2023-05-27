package manager

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/google/uuid"
	"github.com/okian/servo/v2/kv"
	"github.com/okian/servo/v2/lg"
)

func GetParkingAdminPasswordByPhone(ctx context.Context, cr entity.Credentials) (string, error) {
	ParkingAdminPassword, err := repository.GetParkingAdminPasswordByPhone(ctx, cr.Phone)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("error in retrieving ParkingAdmin, %w", err)
	}

	if ParkingAdminPassword != cr.Password {
		return "", ErrUnauthorized
	}

	sessionToken := uuid.NewString()
	lg.Debug(sessionToken)
	kv.Set(ctx, sessionToken, fmt.Sprintf("pAdmin_%s", cr.Phone), time.Second*12000)
	return sessionToken, nil
}
