package prometheus

import "github.com/prometheus/client_golang/prometheus"

func init() {
	prometheus.MustRegister(OnRunCount)
}

var (
	OnRunCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "run_executed",
		Help: "Number of times the controller has checked for cluster state",
	})
)
