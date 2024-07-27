package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Message struct {
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var db *sql.DB

func main() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/messages", createMessageHandler).Methods("POST")

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
