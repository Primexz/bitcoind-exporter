package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TxIndexSynced = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_txindex_synced",
		Help: "The number of blocks in the blockchain",
	})

	TxIndexBestHeight = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_txindex_best_height",
		Help: "The number of headers in the blockchain",
	})
)
