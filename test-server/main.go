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

	fmt.Printf("Database created:%v", db)
	mux := http.NewServeMux()

	server := http.Server{
		Addr: port,
		Handler: mux,
	}

	fmt.Printf("Starting server at port%s", port)
	server.ListenAndServe()
}