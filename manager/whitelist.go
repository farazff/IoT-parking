package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
	"github.com/google/uuid"
	"github.com/okian/servo/v2/lg"
)

func CreateWhitelist(ctx context.Context, Whitelist entity.Whitelist, adminCode uuid.UUID) (int, error) {
	parkingUUID, err := GetParkingUUID(ctx, adminCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return 0, ErrNotFound
		}
		return 0, fmt.Errorf("error in finding parking admin with given id, %w", err)
	}
	id, err := repository.CreateWhitelist(ctx, Whitelist, parkingUUID)
	if err != nil {
		if errors.Is(err, repository.ErrDuplicateEntity) {
			return id, ErrDuplicateEntity
		}
		lg.Error("error during creating Whitelist: %v", err)
		return id, ErrInternalServer
	}
	return id, nil
}

//func GetWhitelists(ctx context.Context, req entity.WhitelistGetReq) ([]entity.Whitelist, error) {
//	parkingId, err := GetParkingUUID(ctx, req.AdminCode)
//	if err != nil {
//		if errors.Is(err, repository.ErrNotFound) {
//			return make([]entity.Whitelist, 0), ErrNotFound
//		}
//		return make([]entity.Whitelist, 0), fmt.Errorf("error in finding parking admin with given id, %w", err)
//	}
//	Whitelists, err := repository.GetWhitelists(ctx, parkingId)
//	if err != nil {
//		return nil, fmt.Errorf("error in retrieving Whitelists, %w", err)
//	}
//	return Whitelists, nil
//}

//func DeleteWhitelist(ctx context.Context, req entity.WhitelistDeleteReq) error {
//	parkingId, err := GetParkingUUID(ctx, req.AdminCode)
//	if err != nil {
//		if errors.Is(err, repository.ErrNotFound) {
//			return ErrNotFound
//		}
//		return fmt.Errorf("error in finding parking admin with given id, %w", err)
//	}
//
//	err = repository.DeleteWhitelist(ctx, parkingId, req.CarTag)
//	if err != nil {
//		if errors.Is(err, repository.ErrNotFound) {
//			return ErrNotFound
//		}
//		return fmt.Errorf("error in finding Whitelist with given id, %w", err)
//	}
//	return nil
//}
