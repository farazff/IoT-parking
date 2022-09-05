package postgres

import (
	"context"
	"fmt"

	"github.com/farazff/IoT-parking/repository"
)

type service struct{}

func (s *service) Name() string {
	return "postgres"
}

func (s *service) Initialize(_ context.Context) error {
	if err := repository.RegisterParking(s); err != nil {
		return fmt.Errorf("error while registring parking repository: %w", err)
	}

	if err := repository.RegisterSystemAdmin(s); err != nil {
		return fmt.Errorf("error while registring system_admin repository: %w", err)
	}
	return nil
}

func (s *service) Finalize() error {
	return nil
}
