package main

import (
	"Lecture9/handlers"
	"Lecture9/internal/user/database"
	"Lecture9/internal/user/repository"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func main() {
	r := gin.Default()

	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "773504ok",
		DB:       0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

	r.Run(":6060")
}
