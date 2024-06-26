package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	TotalConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_total_connections",
		Help: "The number of blocks in the blockchain",
	})

	ConnectionsIn = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_connections_in",
		Help: "The number of headers in the blockchain",
	})

	ConnectionsOut = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_connections_out",
		Help: "The number of headers in the blockchain",
	})

	TotalBytesRecv = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_total_bytes_recv",
		Help: "The number of bytes received",
	})

	TotalBytesSent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_total_bytes_sent",
		Help: "The number of bytes sent",
	})
)
