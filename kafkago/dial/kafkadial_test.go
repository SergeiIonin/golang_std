package dial

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_dialKafka(t *testing.T) {
	pwd, _ := os.Getwd()
	dockerComposeDir := fmt.Sprintf("%s/%s", pwd, "/docker_test/docker-compose.yaml")

	identifier := tc.StackIdentifier("kafka_dial_test")

	compose, err := tc.NewDockerComposeWith(tc.WithStackFiles(dockerComposeDir), identifier)

	compose.Up(context.TODO())

	// cleanup with docker compose down
	t.Cleanup(func() {
		assert.NoError(t, compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal), "compose.Down()")
	})

	if err != nil {
		log.Fatal("Failed to create compose: ", err)
	}

	err = compose.WaitForService("kafka", wait.ForListeningPort("9092")).Up(context.TODO(), tc.Wait(true))

	if err != nil {
		log.Fatal("Failed to wait kafka: ", err)
	}

	brokers := []string{"127.0.0.1:9092"}

	conn, err := kafka.Dial("tcp", brokers[0])
	if err != nil {
		assert.NoError(t, err)
		log.Fatal("failed to dial leader:", err)
	}

	type args struct {
		ctx     context.Context
		timeout time.Duration
		network string
		broker  string
	}
	tests := []struct {
		name    string
		args    args
		want    *kafka.Conn
		wantErr bool
	}{
		{
			"timeout elapsed first",
			args{
				context.Background(),
				3 * time.Second,
				"tcp",
				brokers[0],
			},
			conn,
			true,
		},
		{
			"succesful dial",
			args{
				context.Background(),
				30 * time.Second,
				"tcp",
				brokers[0],
			},
			conn,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dialKafka(tt.args.ctx, tt.args.timeout, tt.args.network, tt.args.broker)
			if (err != nil) && tt.wantErr {
				log.Printf("dialKafka() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				t.Errorf("dialKafka() = %v, want %v", got, tt.want)
			}
		})
	}
}
