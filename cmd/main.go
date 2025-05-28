package main

import (
	"log"
	"net/http"

	"github.com/Robert076/readme-generator/internal/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/readme", api.GetReadme)
	r.Post("/readme", api.PostReadme)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Cannot start HTTP server")
		return
	}
}
