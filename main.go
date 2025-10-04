package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"example.com/pz4-todo/internal/task"
	myMW "example.com/pz4-todo/pkg/middleware"
)

func main() {
	dataPath := os.Getenv("DATA_FILE")
	if dataPath == "" {
		dataPath = filepath.FromSlash("./data/tasks.json")
	}

	repo, err := task.NewRepoWithFile(dataPath)
	if err != nil {
		log.Fatalf("failed to init repo: %v", err)
	}
	h := task.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(myMW.Logger)
	r.Use(myMW.SimpleCORS)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// versioned API
	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/tasks", h.Routes())
	})

	addr := ":8080"
	log.Printf("listening on %s (data file: %s)", addr, dataPath)
	log.Fatal(http.ListenAndServe(addr, r))
}
