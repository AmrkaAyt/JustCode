package routes

import (
	"Lecture8/handlers"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	productGroup := router.Group("/products")
	{
		productGroup.Use(handlers.AuthMiddleware())
		productGroup.POST("/", handlers.CreateProduct)
		productGroup.GET("/", handlers.GetProducts)
		productGroup.GET("/:id", handlers.GetProduct)
		productGroup.PUT("/:id", handlers.UpdateProduct)
		productGroup.DELETE("/:id", handlers.DeleteProduct)
	}
}
