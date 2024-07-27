package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save to DB and send to Kafka (implementation needed)

	w.WriteHeader(http.StatusCreated)
}

func GetMessageStats(w http.ResponseWriter, r *http.Request) {
	// Fetch stats from DB (implementation needed)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Statistics"))
}
