package db

import (
	"context"
	prometheus2 "github.com/okian/servo/v2/monitoring/prometheus"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	metricOpenConn  *prometheus.SummaryVec
	metricInuseConn *prometheus.SummaryVec
	metricIdleConn  *prometheus.SummaryVec
	metricTotalWait *prometheus.SummaryVec
)

func metrics() {
	metricOpenConn = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: "db",
		Name:      "open_conn",
	}, []string{
		"host",
	})
	metricInuseConn = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: "db",
		Name:      "inuse_conn",
	}, []string{
		"host",
	})
	metricIdleConn = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: "db",
		Name:      "idle_conn",
	}, []string{
		"host",
	})
	metricTotalWait = promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: prometheus2.Namespace(),
		Subsystem: "db",
		Name:      "total_wait",
	}, []string{
		"host",
	})

}

func monitor(ctx context.Context, d *sqlx.DB, host string) {

	t := time.Tick(time.Second)
	for {
		select {
		case <-t:
			st := d.Stats()
			metricOpenConn.WithLabelValues(host).Observe(float64(st.OpenConnections))
			metricInuseConn.WithLabelValues(host).Observe(float64(st.InUse))
			metricIdleConn.WithLabelValues(host).Observe(float64(st.Idle))
			metricTotalWait.WithLabelValues(host).Observe(float64(st.WaitDuration.Milliseconds()))
		case <-ctx.Done():
			return
		}
	}
}
