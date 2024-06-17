package zmq

import (
	"context"
	"time"

	"github.com/Primexz/bitcoind-exporter/config"
	prometheus "github.com/Primexz/bitcoind-exporter/prometheus/metrics"
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

	go resetTransactionCount()

	log.WithField("address", address).Info("Listening for zmq messages")

	for {
		// Read envelope
		msg, err := sub.Recv()
		if err != nil {
			log.WithError(err).Fatal("could not receive")
		}

		log.WithField("sequence", msg.Frames[2]).Trace("Transaction received")

		trackTransactionCount()
	}
}

func resetTransactionCount() {
	for {
		prometheus.TransactionsPerSecond.Set(0)
		time.Sleep(time.Second)
	}
}

func trackTransactionCount() {
	prometheus.TransactionsPerSecond.Inc()

}
