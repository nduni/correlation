package app

import (
	"fmt"

	"github.com/nduni/correlation/common/configuration"
)

var Config configuration.Configuration

func LoadConfig() error {
	config, err := configuration.ReadConfig()
	fmt.Println(config)
	if err != nil {
		return err
	}
	Config = config
	return nil
}
