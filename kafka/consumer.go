package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/taturazova/messagio-test/database"
	model "github.com/taturazova/messagio-test/models"
)

// Process a Kafka message
func ProcessMessage(msg model.MessagePayload) {
	database.UpdateMessageStatus(msg.ID, "processed")
}

// Starts a Kafka consumer and processes messages
func StartConsumer() {

	// Get env variables
	brokers := os.Getenv("KAFKA_BROKERS")
	topicName := os.Getenv("KAFKA_TOPIC_NAME")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	if brokers == "" {
		brokers = "localhost:9092" // Default if not set
	}
	if topicName == "" {
		topicName = "default_topic"
	}
	if groupID == "" {
		groupID = "default_group_id"
	}

	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}

	err = consumer.Subscribe(topicName, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("Consumer started")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Fatalf("Failed to read message: %s\n", err)
			continue
		}
		var payload model.MessagePayload
		err = json.Unmarshal(msg.Value, &payload)
		if err != nil {
			log.Fatalf("Failed to unmarshal message payload: %s\n", err)
			continue
		}
		ProcessMessage(payload)
	}
}
