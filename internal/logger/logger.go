package logger

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/log"
)

const (
	envLocal = "local"
)

func SetupLogger(env string) *slog.Logger {
	var logger *slog.Logger

	switch env {
	case envLocal:
		logger = slog.New(log.New(os.Stdout))
	}

	return logger
}

// ErrAttr creates slog attribute with key="error"
func ErrAttr(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
