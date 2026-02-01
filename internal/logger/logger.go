package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
)

const (
	envLocal = "local"
)

// SetupLogger sets default slog logger based on the environment
func SetupLogger(env string) {
	var handler slog.Handler

	switch env {
	case envLocal:
		handler = tint.NewHandler(os.Stdout, &tint.Options{
			Level: slog.LevelDebug,
		})
	}

	slog.SetDefault(slog.New(handler))
}

// ErrAttr creates slog attribute with key="error"
func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}
