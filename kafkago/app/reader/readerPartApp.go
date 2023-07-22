package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SergeiIonin/golang_std/kafkago/readerpartitioned"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddr := "localhost:29092"

	kafkaReaderConfig := &kafka.ReaderConfig{
		Brokers: []string{brokerAddr},
		GroupID: "group_1",
		Topic:   "partitioned_topic_1",
	}

	kafkaReader := readerpartitioned.CreateKafkaReader(kafkaReaderConfig)

	msgs := readerpartitioned.ReadMessages(kafkaReader, context.Background(), 30*time.Second)

	msgsVals := make([]string, len(msgs))
	for i, msg := range msgs {
		msgsVals[i] = string(msg.Value)
	}

	fmt.Printf("msgs received: %v\n", msgsVals)
}
