package model

import (
	"time"
)

type Message struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessagePayload struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
