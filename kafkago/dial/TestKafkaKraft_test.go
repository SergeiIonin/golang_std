package dial

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/segmentio/kafka-go"
	"log"
	"slices"
	"sync"
	"testing"
	"time"
)

func TestKafkaKraft_test(t *testing.T) {
	dockerClient, err := client.NewClientWithOpts(client.WithVersion("1.45"))
	if err != nil {
		log.Printf("error creating docker client: %s", err.Error())
		panic(err)
	}
	container_id, err := CreateKafkaWithKRaftContainer(dockerClient)

	cleanup := func() {
		dockerClient.ContainerRemove(context.Background(), container_id, container.RemoveOptions{Force: true})
	}

	t.Cleanup(cleanup)
	defer cleanup() // fixme

	t.Logf("Container ID: %s", container_id)
	kafkaAddr := "localhost:9092"
	t.Logf("Kafka address: %s", kafkaAddr)

	client := &kafka.Client{
		Addr: kafka.TCP(kafkaAddr),
		//Timeout:   30 * time.Second,
		Transport: nil,
	}

	_, err = client.Heartbeat(context.Background(), &kafka.HeartbeatRequest{
		Addr:            kafka.TCP(kafkaAddr),
		GroupID:         "",
		GenerationID:    0,
		MemberID:        "",
		GroupInstanceID: "",
	})
	if err != nil {
		log.Fatal("failed to heartbeat:", err)
	}
	log.Println("Heartbeat is successful")

	topics := []string{"test_1", "test_2", "test_3"}
	topicsConfigs := make([]kafka.TopicConfig, 0, len(topics))
	for _, topic := range topics {
		conf := kafka.TopicConfig{
			topic,
			1,
			1,
			nil,
			nil,
		}
		topicsConfigs = append(topicsConfigs, conf)
	}
	createTopicsReq := &kafka.CreateTopicsRequest{
		Addr:         kafka.TCP(kafkaAddr),
		Topics:       topicsConfigs,
		ValidateOnly: false,
	}

	_, err = client.CreateTopics(context.Background(), createTopicsReq)
	if err != nil {
		log.Fatal("failed to create topics:", err)
	}

	readerConfig := kafka.ReaderConfig{
		Brokers:  []string{kafkaAddr},
		Topic:    "test_1",
		GroupID:  "test_twks",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	}

	reader := kafka.NewReader(readerConfig)

	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaAddr),
		Topic:    "test_1",
		Balancer: &kafka.LeastBytes{},
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("AAA"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("BBB"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("CCC"),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	msgs := make([]string, 3)

	for i, _ := range msgs {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		msgStr := string(m.Value)
		msgs[i] = msgStr
		log.Printf("message: %s", msgStr)
	}

	log.Printf("msgs: %v", msgs)

	res := slices.Equal([]string{"AAA", "BBB", "CCC"}, msgs)

	if !res {
		t.Fatal("assertion for msgs failed")
	}

}

func CreateKafkaWithKRaftContainer(dockerClient *client.Client) (id string, err error) {
	ctx := context.Background()

	// Define the container configuration
	config := &container.Config{
		Image: "apache/kafka:3.7.0",
		ExposedPorts: nat.PortSet{
			"9092": struct{}{},
		},
		Tty: false,
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"9092": []nat.PortBinding{
				{
					//HostIP: "localhost",
					HostIP:   "0.0.0.0",
					HostPort: "9092",
				},
			},
		},
	}

	// Create the container
	resp, err := dockerClient.ContainerCreate(ctx, config, hostConfig, nil, nil, "kafka")
	if err != nil {
		panic(err)
	}

	// Start the container
	if err := dockerClient.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		panic(err)
	}

	log.Println("WAITING...")

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(7 * time.Second)
		wg.Done()
	}()
	wg.Wait()

	// fixme
	/*isReady := func() bool {
		_, err := dockerClient.Ping(context.Background()) //httpClient.Get("http://localhost:9092")
		if err != nil {
			log.Println("Kafka is not ready yet")
			return false
		}
		log.Println("Kafka is ready")
		return true
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	sleepTime := time.Duration(125) * time.Millisecond
	go func() {
		for {
			if isReady() {
				wg.Done()
				break
			}
			sleepTime = sleepTime * 2
			time.Sleep(sleepTime)
		}
	}()
	wg.Wait()*/

	return resp.ID, nil
}
