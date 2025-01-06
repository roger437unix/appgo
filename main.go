package main

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json: "message"`
}

func main() {
	server := gin.Default()
	server.GET("/", handleRoute)
	server.Run(":8080")
}

func handleRoute(c *gin.Context) {
	response := Response{
		Message: "Hello World!",
	}

	c.JSON(200, response)
}
