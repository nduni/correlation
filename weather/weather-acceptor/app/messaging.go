package app

import (
	"github.com/nduni/correlation/common/messaging"
)

const TOPIC_WEATHER = "weather"

var (
	Senders    map[string]messaging.Sender
	HTTPClient messaging.HTTPRequester
)

func InitMessageBroker() error {
	senders, _, err := messaging.StartMessageBroker(Config.BrokerConnections)
	if err != nil {
		return err
	}
	log.Info().Msgf("message broker initialized")
	Senders = senders

	return nil
}

func InitHTTPClients() {
	HTTPClient = &messaging.RestyClient{Url: weather_api}
}
