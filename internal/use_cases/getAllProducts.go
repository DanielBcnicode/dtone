package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type GetAllProductsDto struct {
	UserId string
	Limit  int
	Page   int
}

type IGetAllProductsUseCase interface {
	Execute(dto GetAllProductsDto) ([]models.Product, error)
}

type GetAllProductsUseCase struct {
	productRepo repositories.ProductRepository
}

func NewGetAllProductsUseCase(productRepo repositories.ProductRepository) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{productRepo: productRepo}
}

func (c *GetAllProductsUseCase) Execute(in GetAllProductsDto) ([]models.Product, error) {
	user, err := c.productRepo.FindAll(in.UserId, in.Limit, in.Page)
	if err != nil {
		return nil, errors.New("the product does not exist")
	}

	return user, nil
}
