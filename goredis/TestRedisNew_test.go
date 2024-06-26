package goredis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
	tcWait "github.com/testcontainers/testcontainers-go/wait"
	"log"
	"time"

	"testing"
)

var (
	ctx            context.Context
	redisClient    *redis.Client
	redisContainer tc.Container
)

func init() {
	ctx = context.Background()

	req := tc.ContainerRequest{
		Image:        "redis:latest",
		Name:         "redis_test",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   tcWait.ForListeningPort("6379/tcp").WithStartupTimeout(60 * time.Second),
	}

	var err error
	redisContainer, err = tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer redisContainer.Terminate(ctx) // Ensure the container is terminated after the test

	host, err := redisContainer.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}
	port, err := redisContainer.MappedPort(ctx, "6379/tcp")
	if err != nil {
		log.Fatal(err)
	}

	redisAddr := fmt.Sprintf("%s:%s", host, port.Port())

	//redisAddr := "localhost:6379"
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}

func TestRedisContainer(t *testing.T) {
	defer func(redisContainer tc.Container, ctx context.Context) {
		err := redisContainer.Terminate(ctx)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}(redisContainer, ctx)

	t.Run("basic KV operations", func(t *testing.T) {
		err := redisClient.Ping(ctx).Err()
		if err != nil {
			t.Fatal(err)
		}

		err = redisClient.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			t.Fatal(err)
		}

		val, err := redisClient.Get(ctx, "key").Result()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Value: %s", val)
		assert.Equal(t, "value", val)
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

		err = redisClient.HSet(ctx, userKey, "userdata", dataBin).Err()
		if err != nil {
			t.Logf("Error at HSet: %s", err)
			t.Fatal(err)
		}

		userDataBin, err := redisClient.HGet(ctx, userKey, "userdata").Result()
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
	t.Run("set and get struct using redis hash", func(t *testing.T) {
		userKey := "user2"

		dataBin, err := json.Marshal(userSchemas)
		if err != nil {
			t.Logf("Error marshalling usersData: %s", err)
			t.Fatal(err)
		}

		err = redisClient.HSet(ctx, userKey, "userSchemas", dataBin).Err()

		if err != nil {
			t.Logf("Error at HSet: %s", err)
			t.Fatal(err)
		}

		userDataBin, err := redisClient.HGet(ctx, userKey, "userSchemas").Result()
		t.Logf("User schemas: %s", userDataBin)
		t.Logf("--------------")
		var userDataFetched []Schema
		err = json.Unmarshal([]byte(userDataBin), &userDataFetched)
		if err != nil {
			t.Logf("Error at HGet: %s", err)
			t.Fatal(err)
		}
		assert.Equal(t, 3, len(userDataFetched))
		t.Logf("Userdata: %v", userDataFetched)
	})

	unmarshalSchema := func(schemaRaw string) Schema {
		var schema Schema
		err := json.Unmarshal([]byte(schemaRaw), &schema)
		if err != nil {
			t.Logf("Error unmarshaling schema: %s", err)
			t.Fatal(err)
		}
		return schema
	}

	t.Run("add structs as separate entries in redis hash", func(t *testing.T) {
		userKey := "user3"
		for _, schema := range userSchemas {
			dataBin, err := json.Marshal(schema)
			if err != nil {
				t.Logf("Error marshalling schema: %s", err)
				t.Fatal(err)
			}
			schemaKey := fmt.Sprintf("%s.%d", schema.Subject, schema.Version)
			err = redisClient.HSet(ctx, userKey, schemaKey, dataBin).Err()
			if err != nil {
				t.Logf("Error at HSet: %s", err)
				t.Fatal(err)
			}
		}

		schemaFooRaw, err := redisClient.HGet(ctx, userKey, "foo.1").Result()
		assert.NoError(t, err)
		schemaBarRaw, err := redisClient.HGet(ctx, userKey, "bar.2").Result()
		assert.NoError(t, err)
		schemaBazRaw, err := redisClient.HGet(ctx, userKey, "baz.3").Result()
		assert.NoError(t, err)

		schemaFoo := unmarshalSchema(schemaFooRaw)
		schemaBar := unmarshalSchema(schemaBarRaw)
		schemaBaz := unmarshalSchema(schemaBazRaw)

		assert.Equal(t, "foo", schemaFoo.Subject)
		assert.Equal(t, 1, schemaFoo.Version)

		assert.Equal(t, "bar", schemaBar.Subject)
		assert.Equal(t, 2, schemaBar.Version)

		assert.Equal(t, "baz", schemaBaz.Subject)
		assert.Equal(t, 3, schemaBaz.Version)
	})
}
