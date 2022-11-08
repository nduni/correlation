package app

import (
	"github.com/Azure/go-amqp"
	"github.com/nduni/correlation/common/messaging/rabbitmq"
)

const TOPIC_CMD_WEATHER = "weather.cmd.1"

var Senders map[string]amqp.Sender

func InitBrokerSenders() error {
	newSenders, err := rabbitmq.StartRabbitmqSenders(Config.BrokerConnections)
	if err != nil {
		return err
	}
	log.Info().Msgf("message broker senders initialized")
	Senders = newSenders
	return nil
}
