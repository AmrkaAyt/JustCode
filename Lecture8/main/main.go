package main

import (
	"Lecture8/db"
	"Lecture8/handlers"
	"Lecture8/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.InitDB()

	r.Use(handlers.AuthMiddleware())

	routes.SetupProductRoutes(r)

	_ = r.Run(":5050")
}
