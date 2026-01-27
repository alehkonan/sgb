package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alehkonan/sgb/packages/storage"
)

type ApiConsumer struct {
	Server *http.ServeMux
	Repo   storage.Storage
}

func New(repo storage.Storage) *ApiConsumer {
	mux := http.NewServeMux()

	return &ApiConsumer{
		Server: mux,
		Repo:   repo,
	}
}

func (api *ApiConsumer) Start() error {
	log.Print("API server is starting...")

	api.Server.HandleFunc("GET /words", api.wordsHandler)

	if err := http.ListenAndServe(":8080", api.Server); err != nil {
		return fmt.Errorf("start server error: %v", err)
	}

	return nil
}

func (api *ApiConsumer) wordsHandler(w http.ResponseWriter, r *http.Request) {
	words, err := api.Repo.GetWords(context.TODO())
	if err != nil {
		log.Printf("[ERROR] words handler. %v", err)
		http.Error(w, "can't get words", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(words)
	if err != nil {
		log.Printf("[ERROR] words handler. %v", err)
		http.Error(w, "can't get words", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
