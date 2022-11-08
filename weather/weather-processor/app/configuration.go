package app

import (
	"fmt"

	"github.com/nduni/correlation/common/configuration"
	"github.com/nduni/correlation/common/logger"
	"github.com/rs/zerolog"
)

var (
	Config configuration.Configuration
	log    *zerolog.Logger = logger.NewPackageLogger("app")
)

func LoadConfig() error {
	config, err := configuration.ReadConfig()
	fmt.Println(config)
	if err != nil {
		return err
	}
	Config = config
	return nil
}
