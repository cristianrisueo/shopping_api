package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// NewLogger returns a zerolog.Logger configured for the current environment.
// In non-release mode it uses a human-readable console output; in release mode it outputs JSON.
func NewLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = time.DateTime

	if os.Getenv("GIN_MODE") != "release" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	return log.Logger
}
