package app

import (
	"context"

	"github.com/nduni/correlation/common/messaging"
)

const TOPIC_WEATHER = "weather"

var Receivers map[string]messaging.Receiver

func StartSubscription() error {
	err := InitMessageBroker()
	if err != nil {
		return err
	}
	startReceiving()
	return nil
}

func InitMessageBroker() error {
	_, receivers, err := messaging.StartMessageBroker(Config.BrokerConnections)
	if err != nil {
		return err
	}
	log.Info().Msgf("message broker initialized")
	Receivers = receivers

	return nil
}

func startReceiving() {
	for topic := range Receivers {
		switch topic {
		case TOPIC_WEATHER:
			Receivers[topic].Receive(context.Background(), ProcessWeather)
		}
	}
}
