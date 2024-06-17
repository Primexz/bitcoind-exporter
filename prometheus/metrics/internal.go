package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ScrapeTime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoind_exporter_scrape_time",
		Help: "The time it took to scrape the data from the bitcoind",
	})
)
