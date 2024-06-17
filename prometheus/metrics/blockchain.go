package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	BlockchainBlocks = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_blockchain_blocks",
		Help: "The number of blocks in the blockchain",
	})

	BlockchainHeaders = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_blockchain_headers",
		Help: "The number of headers in the blockchain",
	})

	BlockchainVerificationProgress = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_blockchain_verification_progress",
		Help: "The verification progress of the blockchain",
	})

	BlockchainSizeOnDisk = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_blockchain_size_on_disk",
		Help: "The size of the blockchain on disk",
	})
)
