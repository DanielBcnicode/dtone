package controller

import (
	"dtonetest/internal/services"
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

// GetOneProducts godoc
// @Sumary        Get one products
// @Description   Get one specific products
// @Tags          product
// @Accept        json
// @Produce       json
// @Param         product_id	path	string	true	"Product ID"
// @Success 200 {object} controller.ProductOutputDto
// @Failure       400
// @Failure       500
// @Security JWT
// @Router /api/v1/products/{product_id} [get]
func (gController *GetOneProductController) Handle(c *gin.Context) {
	productId := c.Param("product_id")

	dtoIn := use_cases.GetOneProductDto{ID: productId}
	user, err := gController.GetOneProductUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gController.modelToOutputDto(user))
}

func (gController *GetOneProductController) modelToOutputDto(m *models.Product) ProductOutputDto {
	return ProductOutputDto{
		ID:          m.ID,
		UserID:      m.UserID,
		Name:        m.Name,
		Description: m.Description,
		File:        m.File,
		Version:     m.Version,
		Price:       services.CoinInt64ToString(m.Price),
	}
}
