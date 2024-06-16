package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sondrefjellving/go-playground/test-server/internal/database"
)

type apiConfig struct {
	db *database.DB
}

func main() {
	port := ":8080"
	path := "database.json"

	db, err := database.NewDB(path)
	if err != nil {
		log.Fatal(err)
	}

	cfg := apiConfig{
		db: db,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/users", cfg.handlerUsersGet)
	mux.HandleFunc("GET /api/users/{userId}", cfg.handlerUsersGetById)
	mux.HandleFunc("POST /api/users", cfg.handlerUsersPost)

	server := http.Server{
		Addr: port,
		Handler: mux,
	}

	fmt.Printf("Starting server at port%s\n", port)
	server.ListenAndServe()
}