package rest

import (
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"

	"github.com/labstack/echo/v4"
	prometheus2 "github.com/okian/servo/v2/monitoring/prometheus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestsCounter *prometheus.CounterVec
	responseTime    *prometheus.HistogramVec
	responseSize    *prometheus.CounterVec
	requestSize     *prometheus.CounterVec
)

const (
	subsystem = "rest"
)

func (s *service) Statictis() {
	requestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: subsystem,
		Name:      "request_total",
	}, []string{
		"path",
		"code",
		"method",
	})

	responseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: subsystem,
		Name:      "response_time",
		Buckets:   prometheus.ExponentialBuckets(0.1, 1.15, 100),
	}, []string{
		"path",
		"code",
		"method",
	})

	responseSize = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: subsystem,
		Name:      "response_size",
	}, []string{
		"path",
		"code",
	})
	requestSize = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: subsystem,
		Name:      "request_size",
	}, []string{
		"path",
		"method",
	})

}

func statictis(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		reqSize := computeApproximateRequestSize(c.Request())
		path := c.Path()
		method := c.Request().Method
		start := time.Now()
		err = next(c)

		if err != nil {
			c.Error(err)
		}

		if err == echo.ErrNotFound && !viper.GetBool(subsystem+".include_not_found") {
			path = "not_registered"
		}

		code := strconv.Itoa(c.Response().Status)
		elapsed := float64(time.Since(start)) / float64(time.Second)
		resSz := float64(c.Response().Size)
		requestsCounter.WithLabelValues(path, code, method).Inc()
		responseTime.WithLabelValues(path, code, method).Observe(elapsed)
		responseSize.WithLabelValues(path, code).Add(resSz)
		requestSize.WithLabelValues(path, method).Add(float64(reqSize))
		return
	}
}

func NetworkWriteSize(path string, code int, size uint64) {
	responseSize.WithLabelValues(path, strconv.Itoa(code)).Add(float64(size))
}

func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
