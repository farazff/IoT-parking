package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
)

type ParkingRepository interface {
	GetParkings(ctx context.Context) ([]entity.Parking, error)
}
