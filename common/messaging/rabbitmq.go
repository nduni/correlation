package messaging

import (
	"context"

	"github.com/Azure/go-amqp"
	"github.com/nduni/correlation/common/configuration"
)

type RabbitMqSender struct {
	rabbitSender func(ctx context.Context, msg *amqp.Message) error
}

type RabbitMqReceiver struct {
	rabbitReceiver *amqp.Receiver
	message        []byte
}

func startRabbitmqBroker(config configuration.BrokerConnection) (map[string]Sender, map[string]Receiver, error) {
	log.Info().Msg("connecting to RabbitMQ broker")
	senders := make(map[string]Sender, len(config.SendingTopics.RabbitmqBroker))
	for _, topicConfig := range config.SendingTopics.RabbitmqBroker {
		newSender, err := StartRabbitmqSenders(topicConfig)
		if err != nil {
			return nil, nil, err
		}
		senders[topicConfig.Name] = newSender
	}

	receivers := make(map[string]Receiver)
	for _, topicConfig := range config.ReceivingTopics.RabbitmqBroker {
		newReceiver, err := StartRabbitmqReceivers(topicConfig)
		if err != nil {
			return nil, nil, err
		}
		receivers[topicConfig.Name] = newReceiver
	}
	log.Info().Msg("succesfully connected to RabbitMQ broker")
	return senders, receivers, nil
}

func StartRabbitmqSenders(topicConfig configuration.Topic) (Sender, error) {
	client, err := amqp.Dial(topicConfig.ConnectionString)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	sender, err := session.NewSender(amqp.LinkTargetAddress("/exchange/Corr/" + topicConfig.Name))
	if err != nil {
		return nil, err
	}
	log.Info().Msgf("sender's connection to '%v' initialized", topicConfig.Name)

	return &RabbitMqSender{
		rabbitSender: sender.Send,
	}, nil
}

func StartRabbitmqReceivers(topicConfig configuration.Topic) (Receiver, error) {
	client, err := amqp.Dial(topicConfig.ConnectionString)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	receiver, err := session.NewReceiver(amqp.LinkSourceAddress("/exchange/Corr/" + topicConfig.Name))
	if err != nil {
		return nil, err
	}
	log.Info().Msgf("receiver's connection to '%v' initialized", topicConfig.Name)

	return &RabbitMqReceiver{
		rabbitReceiver: receiver,
	}, nil
}

func (rb RabbitMqSender) Send(ctx context.Context, msg []byte) error {
	return rb.rabbitSender(ctx, amqp.NewMessage(msg))
}

func (rb RabbitMqReceiver) Receive(ctx context.Context, functionToCall func(context.Context, []byte) error) error {
	go func() error {
		for {
			msg, err := rb.rabbitReceiver.Receive(context.Background())
			if err != nil {
				log.Fatal().Msgf("Reading message from AMQP:", err)
			}
			message := msg.GetData()
			rb.rabbitReceiver.AcceptMessage(ctx, msg)
			log.Info().Msgf("Message received: %s", message)

			err = functionToCall(ctx, message)
			if err != nil {
				log.Error().Msg(err.Error())
			}
		}
	}()
	return nil
}
