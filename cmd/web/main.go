package main

import (
	"fmt"
	"os"

	"sample/internal/api"
	"sample/pkg/models"

	"github.com/go-redis/redis"
)

func main() {
	db := connect()
	defer db.Close()

	app := &api.Application{
		Database: &models.Database{Client: db},
	}

	app.RunServer()
}

func connect() *redis.Client {
	addr := os.Getenv("REDIS_ADDRESS")
	pass := os.Getenv("REDIS_PASS")

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
