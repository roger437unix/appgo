package main

import (
	"github.com/gin-gonic/gin"
	"github.com/roger437unix/appgo/controller"
	"github.com/roger437unix/appgo/db"
	"github.com/roger437unix/appgo/repository"
	"github.com/roger437unix/appgo/usecase"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada de usecases
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
