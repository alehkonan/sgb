package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alehkonan/sgb/packages/storage"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, wrapErr(err)
	}

	if err := db.Ping(); err != nil {
		return nil, wrapErr(err)
	}

	_, err = db.Prepare(`CREATE TABLE IF NOT EXISTS words (ru TEXT, ka TEXT)`)
	if err != nil {
		return nil, wrapErr(err)
	}

	return &Storage{db}, nil
}

func (s *Storage) GetWords(ctx context.Context) ([]storage.Word, error) {
	q := `SELECT (ru, ka) from words`

	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("can't get words: %w", err)
	}

	defer rows.Close()

	var words []storage.Word

	for rows.Next() {
		var word storage.Word

		if err := rows.Scan(&word.Ru, &word.Ka); err != nil {
			return words, err
		}

		words = append(words, word)
	}

	if err = rows.Err(); err != nil {
		return words, err
	}

	return words, nil
}

func (s *Storage) SaveWord(ctx context.Context, word *storage.Word) error {
	q := `INSERT INTO words (ru, ka) VALUES (?, ?)`

	_, err := s.db.ExecContext(ctx, q, word.Ru, word.Ka)
	if err != nil {
		return fmt.Errorf("can't save word: %w", err)
	}

	return nil
}

func wrapErr(err error) error {
	return fmt.Errorf("in sqlite storage: %w", err)
}
