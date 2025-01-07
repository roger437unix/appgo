package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roger437unix/appgo/model"
)

type productController struct {
	// Usercase
}

func NewProductController() productController {
	return productController{}
}

// Para tratar a requisição
func (p *productController) GetProducts(ctx *gin.Context) {
	// Por enquanto apenas testar o controller com dados locais
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20.00,
		},
	}
	ctx.JSON(http.StatusOK, products)
}
