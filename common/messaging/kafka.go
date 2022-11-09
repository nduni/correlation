package messaging

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/nduni/correlation/common/configuration"
)

const partition int = 0

type KafkaSender struct {
	kafkaSender *kafka.Conn
}

type KafkaReceiver struct {
	kafkaReceiver *kafka.Reader
}

func startKafkaBroker(config configuration.BrokerConnection) (map[string]Sender, map[string]Receiver, error) {
	log.Info().Msg("connecting to Kafka broker")

	senders := make(map[string]Sender, len(config.SendingTopics.KafkaBroker))
	for _, topic := range config.SendingTopics.KafkaBroker {
		newSender, err := connectSendersToKafkaCluster(topic)
		if err != nil {
			return nil, nil, err
		}
		senders[topic.Name] = newSender
	}

	receivers := make(map[string]Receiver, len(config.ReceivingTopics.KafkaBroker))
	for _, topic := range config.ReceivingTopics.KafkaBroker {
		newReceiver, err := connectReceiversToKafkaCluster(topic)
		if err != nil {
			return nil, nil, err
		}
		receivers[topic.Name] = newReceiver
	}
	log.Info().Msg("succesfully connected to Kafka broker")

	return senders, receivers, nil
}

func connectSendersToKafkaCluster(topicConfig configuration.Topic) (Sender, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", topicConfig.Brokers[0], topicConfig.Name, partition)
	if err != nil {
		return nil, fmt.Errorf("failed to dial leader: %v", err)
	}
	log.Info().Msgf("sender's connection to topic '%v' initialized", topicConfig.Name)

	return &KafkaSender{
		kafkaSender: conn,
	}, nil
}

func connectReceiversToKafkaCluster(topicConfig configuration.Topic) (Receiver, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   topicConfig.Brokers,
		Topic:     topicConfig.Name,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})

	return &KafkaReceiver{kafkaReceiver: r}, nil
}

func (kr KafkaSender) Send(ctx context.Context, msg []byte) error {
	_, err := kr.kafkaSender.Write(msg)
	return err
}

func (kr KafkaReceiver) Receive(ctx context.Context, functionToCall func(context.Context, []byte) error) error {
	go func() error {
		for {
			msg, err := kr.kafkaReceiver.ReadMessage(context.Background())
			if err != nil {
				log.Error().Msgf(err.Error())
			}
			log.Info().Msgf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))
			err = functionToCall(ctx, msg.Value)
			if err != nil {
				log.Error().Msgf(err.Error())
			}
		}
	}()
	return nil
}
