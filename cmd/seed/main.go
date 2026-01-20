package main

import (
	"context"
	"log"
	repo "sgb/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	ctx := context.Background()
	db, err := gorm.Open(sqlite.Open("tmp/test.db"))
	if err != nil {
		log.Fatalf("Repository create error, %v", err)
	}

	db.AutoMigrate(&repo.Word{})

	words := make([]repo.Word, 2)
	words = append(words,
		repo.Word{Ru: "1", Ka: "1"},
		repo.Word{Ru: "2", Ka: "2"},
	)

	if err := gorm.G[repo.Word](db).CreateInBatches(ctx, &words, 5); err != nil {
		log.Fatalf("Error: %v", err)
	}
	// 	ctx,
	// 	&repository.Word{Ru: "sdcsd", Ka: "sdcd"},
	// ); err != nil {
	// 	log.Fatal("Can not create new words")
	// }

	log.Println("Seed data inserted successfully")
}
