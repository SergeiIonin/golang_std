package main

import (
	"context"
	"log"
	"time"

	"github.com/SergeiIonin/golang_std/kafkago/writerpartitioned"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddr := "localhost:29092"

	client := &kafka.Client{
		Addr:      kafka.TCP(brokerAddr),
		Timeout:   30 * time.Second,
		Transport: nil,
	}

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             "partitioned_topic_1",
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
	}

	req := &kafka.CreateTopicsRequest{
		Addr:         kafka.TCP(brokerAddr),
		Topics:       topicConfigs,
		ValidateOnly: false,
	}

	_, err := client.CreateTopics(context.Background(), req)
	if err != nil {
		log.Printf("could not create topic %v", err)
	}

	kafkaWriterConfig := &kafka.WriterConfig{
		Brokers:  []string{brokerAddr},
		Topic:    "partitioned_topic_1",
		Balancer: &kafka.LeastBytes{},
	}

	kafkaWriter := writerpartitioned.CreateKafkaWriter(kafkaWriterConfig)

	rawMessages := []string{"message 1", "message 2", "message 3", "message 4", "message 5"}

	kafkaMessages := writerpartitioned.GenereateMessagesForPartitionedTopic(rawMessages)

	err = writerpartitioned.WriteMessages(kafkaWriter, kafkaMessages)
	if err != nil {
		log.Printf("[!!!] %v, topic %s", err, kafkaWriter.Topic)
	}
}
