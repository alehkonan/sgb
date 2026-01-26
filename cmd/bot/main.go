package main

import (
	"context"
	"log"
	"os"

	"github.com/alehkonan/sgb/packages/clients/tc"
	"github.com/alehkonan/sgb/packages/consumer/ec"
	"github.com/alehkonan/sgb/packages/processors/tp"
	"github.com/alehkonan/sgb/packages/storage/sqlite"
)

const (
	batchSize = 100
)

func main() {
	tgHost := os.Getenv("TG_HOST")
	if tgHost == "" {
		log.Fatal("TG_HOST environment variable is required")
	}

	tgToken := os.Getenv("TG_TOKEN")
	if tgToken == "" {
		log.Fatal("TG_TOKEN environment variable is required")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is required")
	}

	repo, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatalf("storage open error: %v", err)
	}

	err = repo.Init(context.TODO())
	if err != nil {
		log.Fatalf("storage init error: %v", err)
	}

	processor := tp.New(tc.New(tgHost, tgToken), repo)

	log.Print("Bot is running...")

	consumer := ec.New(&processor, &processor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("Bot is stopped", err)
	}
}
