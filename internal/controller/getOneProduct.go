package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOneProductController struct {
	GetOneProductUseCase use_cases.IGetOneProductUseCase
}

func NewGetOneProductController(
	GetOneProductUseCase use_cases.IGetOneProductUseCase) GetOneProductController {
	return GetOneProductController{
		GetOneProductUseCase: GetOneProductUseCase,
	}
}
func (gController *GetOneProductController) Handle(c *gin.Context) {
	productId := c.Param("product_id")

	dtoIn := use_cases.GetOneProductDto{ID: productId}
	user, err := gController.GetOneProductUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": gController.modelToOutputDto(user)})
}

func (gController *GetOneProductController) modelToOutputDto(m *models.Product) ProductOutputDto {
	return ProductOutputDto{
		ID:          m.ID,
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		File:        m.File,
		Version:     m.Version,
	}
}
