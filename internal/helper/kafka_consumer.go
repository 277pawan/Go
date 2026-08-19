package helper

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func KafkaConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		// GroupID: "order-consumer-group",

		StartOffset: kafka.FirstOffset,
	})
	defer reader.Close()

	fmt.Println("Kafka consumer started. Listening for messages...")

	for {
		message, err := reader.ReadMessage(context.Background())

		if err != nil {
			fmt.Println("Error reading message from Kafka:", err)
		}

		fmt.Println("Message key:", string(message.Key))
		fmt.Println("Received message from Kafka topic:", string(message.Value))
		fmt.Println("Partition", message.Partition)
		fmt.Println("Offset", message.Offset)

		fmt.Println("--------------------------------------------------")
	}
}
