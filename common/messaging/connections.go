package messaging

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nduni/correlation/common/configuration"
	"github.com/nduni/correlation/common/logger"
	"github.com/rs/zerolog"
)

const (
	KAFKA_BROKER    = "kafka"
	RABBITMQ_BROKER = "rabbitmq"
)

var log *zerolog.Logger = logger.NewPackageLogger("rabbitmq")

type Sender interface {
	Send(ctx context.Context, msg []byte) error
}

type Receiver interface {
	Receive(ctx context.Context, functionToCall func(context.Context, []byte) error) error
}

func StartMessageBroker(config configuration.BrokerConnection) (map[string]Sender, map[string]Receiver, error) {
	switch config.MessageBrokerType {
	case KAFKA_BROKER:
		return startMessenger(config, startKafkaBroker)
	case RABBITMQ_BROKER:
		return startMessenger(config, startRabbitmqBroker)
	default:
		return nil, nil, errors.New("no message broker types in config")
	}
}

func startMessenger(config configuration.BrokerConnection, connectionCreator func(connection configuration.BrokerConnection) (map[string]Sender, map[string]Receiver, error)) (map[string]Sender, map[string]Receiver, error) {
	return connectionCreator(config)
}

func SendToBroker[model any](ctx context.Context, senders map[string]Sender, topic string, messageModel model) error {
	message, err := json.Marshal(messageModel)
	if err != nil {
		return err
	}
	sender, ok := senders[topic]
	if !ok {
		return fmt.Errorf("key %v in broker senders map doesn't exist", topic)
	}

	log.Info().Msgf("sending message: %v", string(message))
	err = sender.Send(ctx, message)
	if err != nil {
		return fmt.Errorf("sender error: %v", err)
	}
	log.Info().Msg("message has been sent succesfully")

	return nil
}
