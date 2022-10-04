package configuration

type Configuration struct {
	BrokerConnections BrokerConnection `mapstructure:"broker_connections"`
}

type BrokerConnection struct {
	ReceivingTopics []Topic `mapstructure:"receiving_topics"`
	SendingTopics   []Topic `mapstructure:"sending_topics"`
}

type Topic struct {
	Name             string `mapstructure:"name"`
	ConnectionString string `mapstructure:"connection_string"`
}
