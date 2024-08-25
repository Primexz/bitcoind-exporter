package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type config struct {
	RPCAddress string `env:"RPC_ADDRESS,required"`
	RPCCookieFile string `env:"RPC_COOKIEFILE,required"`

	ZmqAddress string `env:"ZMQ_ADDRESS"`

	FetchInterval int `env:"FETCH_INTERVAL" envDefault:"10"`
	MetricPort    int `env:"METRIC_PORT" envDefault:"3000"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

var (
	log = logrus.WithFields(logrus.Fields{
		"prefix": "config",
	})

	C config
)

func init() {
	loadConfiguration()
}

func loadConfiguration() {
	if config, err := env.ParseAs[config](); err == nil {
		log.Debug("Configuration loaded")
		C = config
	} else {
		log.Fatal(err)
	}

}
