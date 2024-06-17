package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TransactionsPerSecond = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_transactions_per_second",
		Help: "Number of transactions per second",
	})
)
