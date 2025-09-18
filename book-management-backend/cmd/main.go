package main

import (
	"fmt"
	"log"
	"net/http"

	config "book-management/configs"
	database "book-management/pkg/databases"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect DB
	database.ConnectPostgres(cfg)

	// Simple HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Book Management API By Tuan Tran!")
	})

	// Run server
	port := cfg.HTTPPort
	log.Printf("🚀 Server running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
