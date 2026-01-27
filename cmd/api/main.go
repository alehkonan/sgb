package main

import (
	"context"
	"log"
	"os"

	"github.com/alehkonan/sgb/packages/consumer/api"
	"github.com/alehkonan/sgb/packages/storage/sqlite"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is required")
	}

	repo, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatalf("storage open error: %v", err)
	}

	if err := repo.Init(context.TODO()); err != nil {
		log.Fatalf("storage init error: %v", err)
	}

	server := api.New(repo)

	if err = server.Start(); err != nil {
		log.Fatalf("[API] error: %v", err)
	}
}
