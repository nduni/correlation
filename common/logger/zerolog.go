package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var logger *zerolog.Logger = newLogger(os.Stdout)

// provides package level sublogger
func NewPackageLogger(packageName string) *zerolog.Logger {
	log := logger.With().Str("package", packageName).Caller().Logger()
	return &log
}

// creates new logger
func newLogger(w io.Writer) *zerolog.Logger {
	deployment_env := os.Getenv("ENVIRONMENT")
	var logger zerolog.Logger

	if deployment_env == "local" {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatLevel: func(i interface{}) string {
				if i == nil {
					i = ""
				}
				return fmt.Sprintf("\x1b[%dm[%v]\x1b[0m", levelColor(i.(string)), strings.ToUpper(i.(string)))
			},
			FormatMessage: func(i interface{}) string {
				return fmt.Sprintf("| %s |", i)
			},
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
			PartsOrder: []string{
				zerolog.LevelFieldName,
				zerolog.TimestampFieldName,
				zerolog.CallerFieldName,
				zerolog.MessageFieldName,
			},
		}).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Logger()
	} else {
		logger = zerolog.New(w).With().Timestamp().Logger()
	}

	// set log level from env
	logLevel := os.Getenv("LOGLEVEL")
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		logger.Warn().Msg(err.Error())
	}
	logger = logger.Level(level)

	logger.Info().Msgf("Logger initialized on %s level", level.String())

	return &logger
}

// colors definitions
const (
	cReset   = 0
	cRed     = 31
	cGreen   = 32
	cYellow  = 33
	cMagenta = 35
)

// set colors to given log level
func levelColor(level string) int {
	switch level {
	case "debug":
		return cMagenta
	case "info":
		return cGreen
	case "warn":
		return cYellow
	case "error", "fatal", "panic":
		return cRed
	default:
		return cReset
	}
}
