package main

import (
	"Message-Listener-Service/commons"
	"Message-Listener-Service/connector/db"
	"Message-Listener-Service/processor"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	broker = "127.0.0.1:62581"

	consumerGroup = "text-messages-consumer-group"
)

func main() {
	db, err := db.NewDataBaseConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	defer db.Close()

	topic := "text-messages"
	// Create Kafka consumer configuration
	consumerConfig := &kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           consumerGroup,
		"auto.offset.reset":  "earliest",
		"enable.auto.commit": false,
	}

	// Create Kafka consumer
	consumer, err := kafka.NewConsumer(consumerConfig)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	// Subscribe to the topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	// Create Kafka producer configuration
	producerConfig := &kafka.ConfigMap{
		"bootstrap.servers": broker,
	}

	// Create Kafka producer
	producer, err := kafka.NewProducer(producerConfig)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Handle termination signals
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Consume and produce Kafka messages
	for {
		select {
		case <-sigchan:
			fmt.Println("Interrupt signal received. Exiting...")
			return
		default:
			// Consume messages
			msg, err := consumer.ReadMessage(time.Second)
			if err == nil {
				fmt.Printf("Received message: %s\n", string(msg.Value))

				// unmarshal the message
				var kafkaMessage commons.KafkaMessage
				err := json.Unmarshal(msg.Value, &kafkaMessage)
				if err != nil {
					fmt.Printf("Failed to unmarshal message: %v\n", err)
					continue
				}
				trasaction := processor.ExtractTransactionDetails(kafkaMessage.Message)
				fmt.Printf("Transaction details: %v\n", trasaction)

			} else if err.(kafka.Error).Code() != kafka.ErrTimedOut {
				fmt.Printf("Failed to read message: %v\n", err)
			}
		}
	}
}
