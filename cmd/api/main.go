package main

import (
	"log/slog"
	"os"

	"github.com/alehkonan/sgb/internal/config"
	"github.com/alehkonan/sgb/internal/logger"
	"github.com/alehkonan/sgb/packages/consumer/api"
	"github.com/alehkonan/sgb/packages/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()
	logger.SetupLogger(cfg.Env)

	repo, err := sqlite.New(cfg.Env)
	if err != nil {
		slog.Error("failed to init storage", logger.ErrAttr(err))
		os.Exit(1)
	}

	server := api.New(repo)
	if err = server.Start(); err != nil {
		slog.Error("fail to start server", logger.ErrAttr(err))
		os.Exit(1)
	}
}
