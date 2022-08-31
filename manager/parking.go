package manager

import (
	"context"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
)

func GetParkings(ctx context.Context) ([]entity.Parking, error) {
	parkings, err := repository.GetParkings(ctx)
	if err != nil {
		return nil, fmt.Errorf("error in retrieving products, %w", err)
	}
	return parkings, nil
}
