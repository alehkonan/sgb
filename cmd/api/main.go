package main

import (
	"os"

	"github.com/alehkonan/sgb/internal/config"
	"github.com/alehkonan/sgb/internal/logger"
	"github.com/alehkonan/sgb/packages/consumer/api"
	"github.com/alehkonan/sgb/packages/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Warn("DB_PATH is not set")
		os.Exit(1)
	}

	repo, err := sqlite.New(cfg.Env)
	if err != nil {
		log.Error("failed to init storage", logger.ErrAttr(err))
		os.Exit(1)
	}

	server := api.New(repo)
	if err = server.Start(); err != nil {
		log.Error("fail to start server", logger.ErrAttr(err))
		os.Exit(1)
	}
}
