package repository

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
)

type ParkingRepository interface {
	CreateParking(ctx context.Context, parking entity.Parking) (int, error)
	GetParking(ctx context.Context, id int) (entity.Parking, error)
	GetParkings(ctx context.Context) ([]entity.Parking, error)
	UpdateParking(ctx context.Context, parking entity.Parking) error
	DeleteParking(ctx context.Context, id int) error
}
