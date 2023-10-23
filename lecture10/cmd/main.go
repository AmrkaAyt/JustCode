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
	_ "os"
)

func main() {
	r := gin.Default()

	configPath := "config/config.yaml"
	appConfig, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	log.Printf("AppConfig: %+v", appConfig)

	db, err := database.NewDBConnection(configPath)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "773504ok",
		DB:       0,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	userRepository := repository.NewUserRepository(db)

	r.POST("/users/register", func(c *gin.Context) {
		handlers.RegisterHandler(c, userRepository, redisClient)
	})

	r.POST("/users/login", func(c *gin.Context) {
		handlers.LoginHandler(c, userRepository)
	})

	r.GET("/users", func(c *gin.Context) {
		handlers.GetAllUsersHandler(c, userRepository)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		handlers.GetUserByIDHandler(c, userRepository, redisClient)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		handlers.UpdateUserHandler(c, userRepository)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		handlers.DeleteUserHandler(c, userRepository)
	})

	port := appConfig.HttpServer.Port
	log.Printf("Server will run on port: %d", port)
	address := fmt.Sprintf(":%d", port)
	log.Printf("Server will run on address: %s", address)
	r.Run(address)
}
