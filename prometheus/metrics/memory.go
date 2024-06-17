package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MemoryUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_used",
		Help: "The number of blocks in the blockchain",
	})

	MemoryFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_free",
		Help: "The number of headers in the blockchain",
	})

	MemoryTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_total",
		Help: "The best block hash in the blockchain",
	})

	MemoryLocked = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_locked",
		Help: "The difficulty of the blockchain",
	})

	ChunksUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_chunks_used",
		Help: "The time of the blockchain",
	})

	ChunksFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_memory_chunks_free",
		Help: "The median time of the blockchain",
	})
)
