package log

import (
	"github.com/sakari-ai/moirai/log/mode"
)

type Option func(*Logger)

func WithModeFromString(value string) Option {
	return func(logger *Logger) {
		m, err := mode.FromString(value)
		if err != nil {
			panic(err)
		}
		logger.Mode = m
	}
}

func withDevelopmentMode() Option {
	return func(logger *Logger) {
		logger.Mode = mode.Development
	}
}
