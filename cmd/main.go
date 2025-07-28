package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	server := &http.Server{
		Addr:        ":9090",
		Handler:     router,
		ReadTimeout: time.Millisecond * 100,
	}

	log.Fatal(server.ListenAndServe())
}
