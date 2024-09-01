package main

import (
	"log"
	"net/http"

	// internal package imports
	"github.com/SmokierLemur51/data-process/handler"
	"github.com/SmokierLemur51/data-process/middleware"

	// external imports
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	// Load our env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// New Mux !!!
	router := http.NewServeMux()

	// configure routes
	handler.ConfigureRoutes(router)
	
	// create our middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":5000",
		Handler: stack(router),
	}

	log.Println("Starting server on port :5000")
	log.Fatal(server.ListenAndServe())
}
