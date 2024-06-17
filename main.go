package main

import (
	"github.com/Primexz/bitcoind-exporter/config"
	"github.com/Primexz/bitcoind-exporter/fetcher"
	"github.com/Primexz/bitcoind-exporter/prometheus"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat:  "2006/01/02 - 15:04:05",
		FullTimestamp:    true,
		QuoteEmptyFields: true,
		SpacePadding:     45,
	})

	log.SetReportCaller(true)
}

func main() {
	log.Info(config.C.RPCAddress)

	go prometheus.Start()

	fetcher.Start()
}
