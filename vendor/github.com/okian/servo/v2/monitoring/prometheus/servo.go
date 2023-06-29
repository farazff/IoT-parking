package prometheus

import (
	"context"
	"net"
	"net/http"

	"github.com/okian/servo/v2/lg"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
)

func (s *service) Name() string {
	return "monitoring"
}

func checkPort(h, p string) error {
	conn, err := net.Listen("tcp", net.JoinHostPort(h, p))
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func (s *service) Initialize(ctx context.Context) error {
	if !viper.GetBool("monitoring") {
		return nil
	}
	host := viper.GetString("monitoring_host")
	port := viper.GetString("monitoring_port")
	if port == "" {
		port = "9001"
	}

	if err := checkPort(host, port); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	s.server = &http.Server{
		Addr:        net.JoinHostPort(host, port),
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}

	// Run server
	go func() {
		lg.Infof("starting monitoring server on :%s", port)
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			lg.Error(err)
		}
	}()

	return nil
}

func (s *service) Finalize() error {
	if s.server != nil {
		return s.server.Shutdown(context.Background())
	}
	return nil
}
