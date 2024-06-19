package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MiningHashrate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bitcoind_mining_hashrate",
		Help: "Mining hashrate in hashes per second",
	}, []string{"blocks"})
)
