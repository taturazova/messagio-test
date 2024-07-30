package kafka

import (
	"encoding/json"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	model "github.com/taturazova/messagio-test/models"
)

// Kafka producer instance
var Producer *kafka.Producer
var topicName string

// InitializeProducer initializes the Kafka producer
func InitializeProducer() error {
	brokers := os.Getenv("KAFKA_BROKERS")
	topicName := os.Getenv("KAFKA_TOPIC_NAME")

	if brokers == "" {
		brokers = "localhost:9092" // Default if not set
	}
	if topicName == "" {
		topicName = "default_topic" // Default if not set
	}

	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		return err
	}
	return nil
}

// ProduceMessage sends a message to Kafka with its ID
func ProduceMessage(message model.MessagePayload) error {

	payloadBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
		Value:          payloadBytes,
	}

	err = Producer.Produce(kafkaMessage, nil)
	if err != nil {
		return err
	}
	return nil
}

// messagePayload := model.MessagePayload{
// 	ID:      createdID,
// 	Content: message.Content,
// }
// payloadBytes, err := json.Marshal(messagePayload)
// if err != nil {
// 	http.Error(w, "Failed to marshal message payload", http.StatusInternalServerError)
// 	return
// }

// topic := "user_messages"
// kafkaMessage := &kafka.Message{
// 	TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
// 	Value:          payloadBytes,
// }

// err = kafkaProducer.Produce(kafkaMessage, nil)
// if err != nil {
// 	http.Error(w, "Failed to send message to Kafka", http.StatusInternalServerError)
// 	return
// }
