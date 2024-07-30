package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	model "github.com/taturazova/messagio-test/models"
)

var DB *sql.DB

func ConnectDB(host, user, password, dbname string, port int) *sql.DB {
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
	return DB
}

func CreateMessage(msg model.Message) (createdID int, err error) {
	var id int
	err = DB.QueryRow("INSERT INTO messages (content, status) VALUES ($1, $2)", msg.Content, msg.Status).Scan(&id)
	return id, err
}

func UpdateMessageStatus(id int, status string) error {
	_, err := DB.Exec("UPDATE messages SET status = $1 WHERE id = $2", status, id)
	return err
}

func GetMessagesStats() (totalMessages int, processedMessages int, err error) {
	query := `SELECT COUNT(*) AS total_messages, 
            COUNT(CASE WHEN status = 'processed' THEN 1 END) AS processed_messages
        	FROM messages;`
	err = DB.QueryRow(query).Scan(&totalMessages, &processedMessages)
	return totalMessages, processedMessages, err
}
