package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	LoginCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "messenger_logins_total",
		Help: "Total number of successful logins",
	})
	RegisterCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "messenger_registrations_total",
		Help: "Total number of successful registrations",
	})
)

func InitMetrics() {
	prometheus.MustRegister(LoginCounter)
	prometheus.MustRegister(RegisterCounter)
}