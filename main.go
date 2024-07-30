package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/taturazova/messagio-test/api"
	"github.com/taturazova/messagio-test/database"
	"github.com/taturazova/messagio-test/kafka"
)

func main() {
	// Retrieve environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Convert the string to an integer
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatalf("Error converting PORT to integer: %v", err)
	}

	database.ConnectDB(dbHost, dbUser, dbPassword, dbName, port)

	r := api.NewRouter()

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
	kafka.StartConsumer()

}
