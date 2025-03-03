package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type CreateProductDto struct {
	UserID      string
	Name        string
	Description string
	File        string
	Version     string
	Price       int64
}

type ICreateProductUseCase interface {
	Execute(dto CreateProductDto) (*models.Product, error)
}

type CreateProductUseCase struct {
	productRepo repositories.ProductRepository
	userRepo    repositories.UserRepository
}

func NewCreateProductUseCase(
	productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		productRepo: productRepo,
		userRepo:    userRepo,
	}
}

func (c *CreateProductUseCase) Execute(in CreateProductDto) (*models.Product, error) {
	_, err := c.userRepo.FindById(in.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	product := models.Product{
		UserID:      in.UserID,
		Name:        in.Name,
		Description: in.Description,
		File:        in.File,
		Version:     in.Version,
		Price:       in.Price,
	}
	err = c.productRepo.Save(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
