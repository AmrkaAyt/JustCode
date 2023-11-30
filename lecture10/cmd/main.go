// main.go
package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"lecture10/config"
	"lecture10/handlers"
	"lecture10/internal/user/database"
	"lecture10/internal/user/repository"
	"log"
)

func main() {
	r := gin.Default()

	configPath := "config/config.yaml"
	appConfig, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	log.Printf("AppConfig: %+v", appConfig)

	// Use the configuration values directly from appConfig
	db, err := database.NewDBConnection(appConfig.Database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     appConfig.Redis.Addr,
		Password: appConfig.Redis.Password,
		DB:       appConfig.Redis.DB,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	// Your route handlers...

	port := appConfig.HttpServer.Port
	log.Printf("Server will run on port: %d", port)
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server will run on address: %s", address)
	r.Run(address)
}
