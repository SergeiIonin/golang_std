package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestKafkaConnection(t *testing.T) { // seems to be flacky
	pwd, _ := os.Getwd()
	dockerComposeDir := fmt.Sprintf("%s/%s", pwd, "/docker_test/docker-compose.yaml")

	compose, err := tc.NewDockerCompose(dockerComposeDir)

	compose.Up(context.TODO())

	if err != nil {
		log.Fatal("Failed to create compose: ", err)
	}

	projects := compose.Services()
	for _, project := range projects {
		log.Println("project :", project)
	}
	assert.NotEmpty(t, projects, "compose.Services()")

	topicConfigs := []kafka.TopicConfig{
		kafka.TopicConfig{
			Topic:             "topic1",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
		kafka.TopicConfig{
			Topic:             "topic2",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = compose.WaitForService("kafka", wait.ForListeningPort("9092")).Up(context.TODO(), tc.Wait(true))

	if err != nil {
		log.Fatal("Failed to wait kafka: ", err)
	}

	brokers := []string{"127.0.0.1:9092"}

	conn, err := kafka.DialLeader(context.Background(), "tcp", brokers[0], "timewindows", 0)
	if err != nil {
		assert.NoError(t, err)
		log.Fatal("failed to dial leader:", err)
	}

	conn.CreateTopics(topicConfigs...)

	container, err := compose.ServiceContainer(context.TODO(), "kafka")
	assert.NoError(t, err)

	ip, err := container.ContainerIP(context.TODO())

	assert.NoError(t, err)
	log.Println(ip)
}
