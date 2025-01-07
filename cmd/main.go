package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roger437unix/appgo/controller"
)

func main() {
	server := gin.Default()

	// Camada de controllers
	ProductController := controller.NewProductController()
	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
