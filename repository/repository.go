package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
)

type ParkingRepository interface {
	GetParking(ctx context.Context, id int) (entity.Parking, error)
	GetParkings(ctx context.Context) ([]entity.Parking, error)
	DeleteParking(ctx context.Context, id int) error
}
