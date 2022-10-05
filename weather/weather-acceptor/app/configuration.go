package app

import (
	"github.com/nduni/correlation/common/configuration"
)

var Config configuration.Configuration

func LoadConfig() error {
	config, err := configuration.ReadConfig()
	if err != nil {
		return err
	}
	Config = config
	return nil
}
