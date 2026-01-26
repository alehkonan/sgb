package storage

import "context"

type Storage interface {
	GetWords(ctx context.Context) ([]Word, error)
	SaveWord(ctx context.Context, word *Word) error
}

type Word struct {
	Ru string `db:"ru" json:"russianWord"`
	Ka string `db:"ka" json:"georgianWord"`
}
