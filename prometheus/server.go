package prometheus

import (
	"net/http"
	"strconv"

	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"prefix": "prometheus",
	})
)

func Start() {
	port := strconv.Itoa(config.C.MetricPort)

	log.WithField("port", port).Info("Starting Prometheus metrics server.")

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}