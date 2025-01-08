package usecase

import (
	"github.com/roger437unix/appgo/model"
	"github.com/roger437unix/appgo/repository"
)

type ProductUsecase struct {
	// Repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
