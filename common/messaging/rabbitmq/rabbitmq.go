package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/go-amqp"
	"github.com/nduni/correlation/common/configuration"
)

func StartRabbitmqSenders(config configuration.BrokerConnection) (map[string]amqp.Sender, error) {
	senders := make(map[string]amqp.Sender, len(config.SendingTopics))
	if len(config.SendingTopics) == 0 {
		return senders, errors.New("no senders configuration")
	}
	for _, brockerConfig := range config.SendingTopics {
		client, err := amqp.Dial(brockerConfig.ConnectionString)
		if err != nil {
			return senders, err
		}
		session, err := client.NewSession()
		if err != nil {
			return senders, err
		}
		sender, err := session.NewSender(amqp.LinkTargetAddress("/exchange/Corr/" + brockerConfig.Name))
		if err != nil {
			return senders, err
		}
		senders[brockerConfig.Name] = *sender
		fmt.Printf("sender's connection to '%v' initialized\n", brockerConfig.Name)
	}
	return senders, nil
}

func SendToBroker[model any](ctx context.Context, senders map[string]amqp.Sender, topic string, messageModel model) error {
	message, err := json.Marshal(messageModel)
	if err != nil {
		return err
	}
	sender, ok := senders[topic]
	if !ok {
		return fmt.Errorf("key %v in broker senders map doesn't exist", topic)
	}
	err = sender.Send(ctx, amqp.NewMessage(message))
	if err != nil {
		return fmt.Errorf("sender error: %v", err)
	}
	return nil
}

func StartRabbitmqReceivers(config configuration.BrokerConnection) (map[string]amqp.Receiver, error) {
	receivers := make(map[string]amqp.Receiver, len(config.ReceivingTopics))
	if len(config.ReceivingTopics) == 0 {
		return receivers, errors.New("no receivers configuration")
	}
	for _, brockerConfig := range config.ReceivingTopics {
		client, err := amqp.Dial(brockerConfig.ConnectionString)
		if err != nil {
			return receivers, err
		}
		session, err := client.NewSession()
		if err != nil {
			return receivers, err
		}
		receiver, err := session.NewReceiver(amqp.LinkSourceAddress("/exchange/Corr/" + brockerConfig.Name))
		if err != nil {
			return receivers, err
		}
		receivers[brockerConfig.Name] = *receiver
		fmt.Printf("receiver's connection to '%v' initialized\n", brockerConfig.Name)
	}
	return receivers, nil
}

func ReceiveFromBroker(ctx context.Context, topic string, receiver amqp.Receiver, messageHandler func(ctx context.Context) error) error {
	go func() error {
		for {
			msg, err := receiver.Receive(context.Background())
			if err != nil {
				log.Fatal("Reading message from AMQP:", err)
			}
			message := msg.GetData()
			receiver.AcceptMessage(ctx, msg)

			fmt.Printf("Message received: %s\n", message)
		}
	}()
	return nil
}
