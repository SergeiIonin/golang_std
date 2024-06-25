package goredis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"time"

	"testing"
)

func TestRedisContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "goredis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp").WithStartupTimeout(60 * time.Second),
	}

	redisContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer redisContainer.Terminate(ctx) // Ensure the container is terminated after the test

	host, err := redisContainer.Host(ctx)
	if err != nil {
		t.Fatal(err)
	}
	port, err := redisContainer.MappedPort(ctx, "6379/tcp")
	if err != nil {
		t.Fatal(err)
	}

	redisAddr := fmt.Sprintf("%s:%s", host, port.Port())
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	t.Run("basic KV operations", func(t *testing.T) {
		err := rdb.Ping(ctx).Err()
		if err != nil {
			t.Fatal(err)
		}

		err = rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			t.Fatal(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Value: %s", val)
		//assert.Equal(t, "value", val)
	})

	// an equivalent of redis-cli commands
	// HSET user1 userdata '{"id": "1", "username": "bob"}'
	// HGET user1 userdata
	t.Run("set and get struct using redis hash", func(t *testing.T) {
		type UserData struct {
			ID       string `json:"id"`
			Username string `json:"username"`
		}

		userData := UserData{
			ID:       "1",
			Username: "bob",
		}

		userKey := "user1"

		dataBin, err := json.Marshal(userData)
		if err != nil {
			t.Logf("Error marshalling userData: %s", err)
			t.Fatal(err)
		}

		err = rdb.HSet(ctx, userKey, "userdata", dataBin).Err()
		if err != nil {
			t.Logf("Error at HSet: %s", err)
			t.Fatal(err)
		}

		userDataBin, err := rdb.HGet(ctx, userKey, "userdata").Result()
		var userDataFetched UserData
		err = json.Unmarshal([]byte(userDataBin), &userDataFetched)
		if err != nil {
			t.Logf("Error at HGet: %s", err)
			t.Fatal(err)
		}
		t.Logf("Userdata: %s", userDataFetched)
	})

	// an equivalent of redis-cli commands
	// HSET user2 userSchemas '[{\"subject\":\"foo\",\"version\":1,\"id\":2,\"fields\":[\"firstName\",\"lastName\"]},{\"subject\":\"bar\",\"version\":2,\"id\":3,\"fields\":[\"userName\",\"deposit\"]},{\"subject\":\"baz\",\"version\":3,\"id\":4,\"fields\":[\"email\",\"phone\"]}]'
	// HGET user2 userSchemas
	t.Run("set and get struct using redis hash", func(t *testing.T) {
		type Schema struct {
			Subject string   `json:"subject"`
			Version int      `json:"version"`
			ID      int      `json:"id"`
			Fields  []string `json:"fields"`
		}
		userSchemas := []Schema{
			{
				Subject: "foo",
				Version: 1,
				ID:      2,
				Fields:  []string{"firstName", "lastName"},
			},
			{
				Subject: "bar",
				Version: 2,
				ID:      3,
				Fields:  []string{"userName", "deposit"},
			},
			{
				Subject: "baz",
				Version: 3,
				ID:      4,
				Fields:  []string{"email", "phone"},
			},
		}

		userKey := "user2"

		dataBin, err := json.Marshal(userSchemas)
		if err != nil {
			t.Logf("Error marshalling usersData: %s", err)
			t.Fatal(err)
		}

		err = rdb.HSet(ctx, userKey, "userSchemas", dataBin).Err()

		if err != nil {
			t.Logf("Error at HSet: %s", err)
			t.Fatal(err)
		}

		userDataBin, err := rdb.HGet(ctx, userKey, "userSchemas").Result()
		t.Logf("User schemas: %s", userDataBin)
		t.Logf("--------------")
		var userDataFetched []Schema
		err = json.Unmarshal([]byte(userDataBin), &userDataFetched)
		if err != nil {
			t.Logf("Error at HGet: %s", err)
			t.Fatal(err)
		}
		t.Logf("Userdata: %v", userDataFetched)
	})
}
