package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
	"os"
	"path/filepath"
)

type UploadProductDto struct {
	UserID    string
	ProductID string
	File      string
}

type IUploadProductUseCase interface {
	Execute(dto UploadProductDto) (*models.Product, error)
}

type UploadProductUseCase struct {
	productRepo        repositories.ProductRepository
	userRepo           repositories.UserRepository
	fileRepositoryPath string
}

func NewUploadProductUseCase(
	productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository,
	fileRepositoryPath string,
) *UploadProductUseCase {
	return &UploadProductUseCase{
		productRepo:        productRepo,
		userRepo:           userRepo,
		fileRepositoryPath: fileRepositoryPath,
	}
}

func (c *UploadProductUseCase) Execute(in UploadProductDto) (*models.Product, error) {
	product, err := c.productRepo.FindById(in.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	if len(product.File) > 0 && product.File != in.File {
		filePath := filepath.Join(c.fileRepositoryPath, product.File)
		err = os.Remove(filePath)
		if err != nil {
			return nil, err
		}
	}

	product.File = in.File
	err = c.productRepo.Save(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
