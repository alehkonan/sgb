package main

import (
	"context"
	"log"
	"os"

	"github.com/alehkonan/sgb/packages/storage"
	"github.com/alehkonan/sgb/packages/storage/sqlite"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is required")
	}

	repo, err := sqlite.New(dbPath)
	if err != nil {
		log.Fatalf("can't open the storage: %v", err)
	}

	words := make([]storage.Word, 2)
	words = append(words,
		storage.Word{Ru: "1", Ka: "1"},
		storage.Word{Ru: "2", Ka: "2"},
	)

	for _, word := range words {
		err = repo.SaveWord(context.TODO(), &word)
		if err != nil {
			log.Printf("Word %s was not saved. Reason: %v", word.Ru, err)
			continue
		}
	}

	log.Println("Seed data inserted successfully")
}
