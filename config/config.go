package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

type config struct {
	RPCAddress string `env:"RPC_ADDRESS,required"`
	RPCUser    string `env:"RPC_USER,required"`
	RPCPass    string `env:"RPC_PASS,required"`
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
