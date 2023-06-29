package lg

import (
	"context"
)

func (s *service) Name() string {
	return "zap_logger"
}

func (s *service) Initialize(_ context.Context) error {
	if err := s.setup(); err != nil {
		return err
	}
	Register(s)
	return nil
}

func (s *service) Finalize() error {
	return s.logger.Sync()
}

func (s *service) Healthy(_ context.Context) (interface{}, error) {
	return nil, s.logger.Sync()
}

func (s *service) Ready(_ context.Context) (interface{}, error) {
	return nil, s.logger.Sync()
}
