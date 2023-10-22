package main

import (
	"Lecture9/handlers"
	"Lecture9/internal/user/database"
	"Lecture9/internal/user/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := database.NewDBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	r.POST("/users/register", func(c *gin.Context) {
		handlers.RegisterHandler(c, userRepository)

	})

	r.POST("/users/login", func(c *gin.Context) {
		handlers.LoginHandler(c, userRepository)
	})

	r.GET("/users", func(c *gin.Context) {
		handlers.GetAllUsersHandler(c, userRepository)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		handlers.GetUserByIDHandler(c, userRepository)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		handlers.UpdateUserHandler(c, userRepository)
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		handlers.DeleteUserHandler(c, userRepository)
	})

	r.Run(":6060")
}
