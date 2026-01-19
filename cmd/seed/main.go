package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:repository/test.db")
	if err != nil {
		log.Fatalf("Database connection error: , %v", err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS words (id INTEGER PRIMARY KEY AUTOINCREMENT, ru TEXT NOT NULL, ka TEXT NOT NULL)")
	if err != nil {
		log.Fatalf("Failed to create a table: %v", err)
	}

	_, err = db.Exec("DELETE FROM words")
	if err != nil {
		log.Fatalf("Failed to delete data: %v", err)
	}

	_, err = db.Exec("INSERT INTO words (ru, ka) VALUES ('светлячок', 'ციცინათელა')")
	if err != nil {
		log.Fatalf("Failed to insert a row: %v", err)
	}

	log.Println("Seed data inserted successfully")
}
