package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const fileprefix = "config_"

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

func ReadConfig() (Configuration, error) {
	var config Configuration
	deployment_env := os.Getenv("DEPLOYMENT_ENV")
	viper.SetConfigName(fileprefix + deployment_env) // name of config file (without extension)
	viper.SetConfigType("yaml")                      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./resources/config")
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	fmt.Println("Load configuration for ", deployment_env)
	return config, nil
}
