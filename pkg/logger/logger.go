package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() (*zerolog.Logger, error) {
	zerolog.TimeFieldFormat = time.RFC3339

	core := zerolog.NewConsoleWriter()

	core.Out = io.MultiWriter(os.Stdout)
	core.TimeFormat = "Jan 2 15:04:05"

	log.Logger = zerolog.New(core).With().Logger()
	return &log.Logger, nil
}
