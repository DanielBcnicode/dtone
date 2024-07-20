package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProductInputDto struct {
	UserID      string `json:"user_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	File        string `json:"-"`
	Version     string `json:"version" binding:"required"`
}

type ProductOutputDto struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	File        string `json:"file"`
	Version     string `json:"version"`
}

type CreateProductController struct {
	CreateProductUseCase use_cases.ICreateProductUseCase
	FileRepository       string
}

func NewCreateProductController(
	CreateUseCase use_cases.ICreateProductUseCase,
	fileRepositoryPath string,
) CreateProductController {
	return CreateProductController{
		CreateProductUseCase: CreateUseCase,
		FileRepository:       fileRepositoryPath,
	}
}

func (cController *CreateProductController) Handle(c *gin.Context) {
	var input CreateProductInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoIn := use_cases.CreateProductDto{
		UserID:      input.UserID,
		Name:        input.Name,
		Description: input.Description,
		File:        "",
		Version:     input.Version,
	}
	product, err := cController.CreateProductUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": cController.modelToOutputDto(product)})
}

func (cController *CreateProductController) modelToOutputDto(m *models.Product) ProductOutputDto {
	return ProductOutputDto{
		ID:          m.ID,
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		File:        m.File,
		Version:     m.Version,
	}
}
