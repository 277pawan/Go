package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func KafkaProducer() {

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
	})

	defer writer.Close()

	message := kafka.Message{Key: []byte("order_102"),
		Value: []byte(
			`{"order_id": 110, "user_id": 2, "product_id": 100, "quantity": 5}`),
	}

	err := writer.WriteMessages(context.Background(), message)

	if err != nil {
		log.Fatal("Failed to write message to Kafka:", err)
	}

	fmt.Println("Message sent to kafka topic successfully")

}
