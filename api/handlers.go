package api

import (
	"encoding/json"
	"net/http"

	"github.com/taturazova/messagio-test/database"
	kafka "github.com/taturazova/messagio-test/kafka"
	model "github.com/taturazova/messagio-test/models"
)

type Stats struct {
	TotalMessages     int `json:"total_messages"`
	ProcessedMessages int `json:"processed_messages"`
}

type Response struct {
	Stats Stats `json:"stats"`
}

func init() {
	kafka.InitializeProducer()
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message model.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save to DB and send to Kafka

	// Saving to DB
	createdID, err := database.CreateMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Producing to Kafka
	messagePayload := model.MessagePayload{
		ID:      createdID,
		Content: message.Content,
	}

	err = kafka.ProduceMessage(messagePayload)
	if err != nil {
		http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "Message saved"})
}

func GetMessageStats(w http.ResponseWriter, r *http.Request) {
	// Fetch stats from DB
	totalMessages, processedMessages, err := database.GetMessagesStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := Response{
		Stats: Stats{
			TotalMessages:     totalMessages,
			ProcessedMessages: processedMessages,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
