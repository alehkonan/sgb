package repository

import (
	"context"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	// TODO use normal DB
	db, err := gorm.Open(sqlite.Open("tmp/test.db"))
	if err != nil {
		log.Fatalf("Repository create error, %v", err)
	}

	db.AutoMigrate(&Word{})

	return &Repository{db}
}

func (r *Repository) GetWords() ([]Word, error) {
	ctx := context.Background()
	words, err := gorm.G[Word](r.db).Find(ctx)

	return words, err
}
