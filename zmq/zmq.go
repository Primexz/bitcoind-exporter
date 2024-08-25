package zmq

import (
	"context"

	"github.com/Pxrzival/bitcoind-exporter/config"
	prometheus "github.com/Pxrzival/bitcoind-exporter/prometheus/metrics"
	"github.com/go-zeromq/zmq4"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"prefix": "zmq",
	})
)

func Start() {
	address := config.C.ZmqAddress
	if address == "" {
		log.Debug("Zmq address not set, skipping zmq listener")
		return
	}

	sub := zmq4.NewSub(context.Background())
	defer sub.Close()

	err := sub.Dial("tcp://" + address)
	if err != nil {
		log.WithError(err).Fatal("could not dial")
	}

	err = sub.SetOption(zmq4.OptionSubscribe, "rawtx")
	if err != nil {
		log.WithError(err).Fatal("could not set option")
	}

	log.WithField("address", address).Info("Listening for zmq messages")

	for {
		// Read envelope
		msg, err := sub.Recv()
		if err != nil {
			log.WithError(err).Fatal("could not receive")
		}

		transaction := string(msg.Frames[1])
		log.WithField("transaction", transaction).Debug("Received transaction")

		prometheus.TransactionsPerSecond.Inc()
	}
}
