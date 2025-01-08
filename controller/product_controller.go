package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roger437unix/appgo/usecase"
)

type productController struct {
	// Usercase
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

// Para tratar a requisição
func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
