package postgres

import (
	"context"
	"github.com/farazff/IoT-parking/entity"
	"github.com/okian/servo/v2/db"
)

const (
	getParkingsQuery = `SELECT id, name, address, phone, enabled, created_at, updated_at, deleted_at FROM parkings WHERE deleted_at is NULL`
)

func (s *service) GetParkings(ctx context.Context) ([]entity.Parking, error) {
	var ps []Parking
	err := db.Select(ctx, &ps, getParkingsQuery)
	if err != nil {
		return nil, err
	}

	res := make([]entity.Parking, 0)
	for i := range ps {
		res = append(res, ps[i])
	}
	return res, nil
}
