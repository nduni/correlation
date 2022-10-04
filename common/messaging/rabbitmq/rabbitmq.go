package rabbitmq

import (
	"github.com/Azure/go-amqp"
	"github.com/nduni/correlation/common/configuration"
)

func StartRabbitmqSenders(config configuration.BrokerConnection) (map[string]*amqp.Sender, error) {
	senders := make(map[string]*amqp.Sender, len(config.SendingTopics))
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
		senders[brockerConfig.Name] = sender
	}
	return senders, nil
}
