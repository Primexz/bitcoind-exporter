package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SmartFee = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "bitcoind_smart_fee",
		Help: "The estimate fee rate in satoshis per byte",
	}, []string{"blocks"})
)
