package app

import (
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
	if err != nil {
		return err
	}
	Config = config
	return nil
}
