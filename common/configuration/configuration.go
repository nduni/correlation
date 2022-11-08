package configuration

import (
	"os"

	"github.com/nduni/correlation/common/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

const fileprefix = "config_"

var log *zerolog.Logger = logger.NewPackageLogger("configuration")

type Configuration struct {
	BrokerConnections BrokerConnection `mapstructure:"broker_connections"`
	Db                DB
}

type BrokerConnection struct {
	ReceivingTopics []Topic `mapstructure:"receiving_topics"`
	SendingTopics   []Topic `mapstructure:"sending_topics"`
}

type Topic struct {
	Name             string `mapstructure:"name"`
	ConnectionString string `mapstructure:"connection_string"`
}

type DB struct {
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_SSL      string `mapstructure:"DB_SSL"`
}

func ReadConfig() (Configuration, error) {
	var config Configuration
	deployment_env := os.Getenv("ENVIRONMENT")
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
	log.Info().Msgf("Load configuration for %v", deployment_env)
	return config, nil
}
