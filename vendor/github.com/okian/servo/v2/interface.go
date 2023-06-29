package servo

import "context"

type Service interface {
	Name() string
	Initialize(context.Context) error
	Finalize() error
}

type Readiness interface {
	Ready(context.Context) (interface{}, error)
}

type Healthiness interface {
	Healthy(context.Context) (interface{}, error)
}
