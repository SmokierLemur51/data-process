package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/SmokierLemur51/lbp-hit-my-line/db"

	"github.com/jmoiron/sqlx"
)

type Handler struct {
	DB *sqlx.DB
}

func ConfigureRoutes(router *http.ServeMux) {
	// Handler first
	h := NewHandler()

	// Test routes to delete
	router.HandleFunc("GET /cmd/{command}", h.CommandHandler)

	// Application routes
	router.HandleFunc("GET /", h.IndexHandler)
	router.HandleFunc("GET /example/{value}", h.ExampleHandler)
}

func NewHandler() *Handler {
	return &Handler{
		DB: ConnectDB(),
	}
}

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect(
		"sqlite3",
		fmt.Sprintf("db/instance/%s.db", os.Getenv("SQLITE_DB")),
	)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
