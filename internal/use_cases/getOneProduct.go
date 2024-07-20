package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type GetOneProductDto struct {
	ID string
}

type IGetOneProductUseCase interface {
	Execute(dto GetOneProductDto) (*models.Product, error)
}

type GetOneProductUseCase struct {
	productRepo repositories.ProductRepository
}

func NewGetOneProductUseCase(productRepo repositories.ProductRepository) *GetOneProductUseCase {
	return &GetOneProductUseCase{productRepo: productRepo}
}

func (c *GetOneProductUseCase) Execute(in GetOneProductDto) (*models.Product, error) {
	user, err := c.productRepo.FindById(in.ID)
	if err != nil {
		return nil, errors.New("the product does not exist")
	}

	return &user, nil
}
