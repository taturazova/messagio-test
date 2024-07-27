package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func ConnectDB(host, user, password, dbname string, port int) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to the database")
}
