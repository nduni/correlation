package app

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/nduni/correlation/common/messaging/rabbitmq"
)

const TOPIC_CMD_WEATHER = "weather.cmd.1"

var Receivers map[string]amqp.Receiver

func StartSubscription() error {
	err := InitBrokerReceivers()
	if err != nil {
		return err
	}
	startReceiving()
	return nil
}

func InitBrokerReceivers() error {
	newReceivers, err := rabbitmq.StartRabbitmqReceivers(Config.BrokerConnections)
	if err != nil {
		return err
	}
	log.Info().Msgf("message broker receivers initialized")
	Receivers = newReceivers

	return nil
}

func startReceiving() {
	for topic := range Receivers {
		switch topic {
		case TOPIC_CMD_WEATHER:
			rabbitmq.ReceiveFromBroker(context.Background(), topic, Receivers[topic], ProcessWeather)
		}
	}
}
