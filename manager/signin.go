package manager

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"

	"github.com/farazff/IoT-parking/entity"
	"github.com/farazff/IoT-parking/repository"
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

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := &entity.CustomClaims{
		Phone: cr.Phone,
		Type:  "parking",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(), // Token expires in 5 minutes
		},
	}

	token.Claims = claims

	// Generate encoded token and send it as response
	t, err := token.SignedString(entity.SecretKey)
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetSystemAdminPasswordByPhone(ctx context.Context, cr entity.Credentials) (string, error) {
	systemAdminPassword, err := repository.GetSystemAdminPasswordByPhone(ctx, cr.Phone)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("error in retrieving SystemAdmin, %w", err)
	}

	if systemAdminPassword != cr.Password {
		return "", ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := &entity.CustomClaims{
		Phone: cr.Phone,
		Type:  "system",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * -5).Unix(), // Token expires in 5 minutes
		},
	}

	token.Claims = claims

	// Generate encoded token and send it as response
	t, err := token.SignedString(entity.SecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func GetUserPasswordByPhone(ctx context.Context, cr entity.Credentials) (string, error) {
	userPassword, err := repository.GetUserPasswordByPhone(ctx, cr.Phone)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}
		return "", fmt.Errorf("error in retrieving user, %w", err)
	}

	if userPassword != cr.Password {
		return "", ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := &entity.CustomClaims{
		Phone: cr.Phone,
		Type:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(), // Token expires in 5 minutes
		},
	}

	token.Claims = claims

	// Generate encoded token and send it as response
	t, err := token.SignedString(entity.SecretKey)
	if err != nil {
		return "", err
	}

	return t, nil
}
