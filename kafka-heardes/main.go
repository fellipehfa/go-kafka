package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func main() {

	//PRODUCER
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "quickstart",
	}

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("menssagem"),
		Headers: []protocol.Header{
			{
				Key:   "session",
				Value: []byte("123"),
			},
		},
	})

	if err != nil {
		log.Fatal("cannot write message", err)
	}

	//CONSUMER
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group",
		Topic:    "quickstart",
		MinBytes: 0,    // 0KB
		MaxBytes: 10e6, // 10MB
	})

	for i := 0; i < 1; i++ {
		message, err := reader.ReadMessage(context.Background())

		for _, val := range message.Headers {
			if val.Key == "session" && string(val.Value) == "123" {
				fmt.Println(string("sessao correta"))
			}
		}

		if err != nil {
			log.Fatal("cannot read message", err)
			reader.Close()
		}

		fmt.Print("receive a message: ", string(message.Value))
	}

	reader.Close()
}
