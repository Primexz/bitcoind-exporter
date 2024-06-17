package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MempoolUsage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_mempool_usage",
		Help: "Total memory usage for the mempool in bytes",
	})

	MempoolMax = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_mempool_max",
		Help: "Maximum memory usage for the mempool in bytes",
	})
)
