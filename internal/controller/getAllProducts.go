package controller

import (
	"dtonetest/internal/services"
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GetAllProductsController struct {
	GetAllProductsUseCase use_cases.IGetAllProductsUseCase
}

func NewGetAllProductsController(
	GetAllProductsUseCase use_cases.IGetAllProductsUseCase) GetAllProductsController {
	return GetAllProductsController{
		GetAllProductsUseCase: GetAllProductsUseCase,
	}
}

// GetAllProducts godoc
// @Sumary        Get all products
// @Description   Get all products
// @Tags          product
// @Accept        json
// @Produce       json
// @Param         user_id	query	string	true	"User ID"
// @Param         page	query	int	true	"Page"
// @Param         limit	query	int	true	"Limit"
// @Success 200 {array} controller.ProductOutputDto
// @Failure       400
// @Failure       500
// @Security JWT
// @Router /api/v1/products [get]
func (gController *GetAllProductsController) Handle(c *gin.Context) {
	userId := c.Query("user_id")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	dtoIn := use_cases.GetAllProductsDto{
		UserId: userId,
		Limit:  limit,
		Page:   page,
	}
	users, err := gController.GetAllProductsUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gController.modelToOutputDto(users))
}

func (gController *GetAllProductsController) modelToOutputDto(m []models.Product) []ProductOutputDto {
	result := make([]ProductOutputDto, len(m))
	for i := 0; i < len(m); i++ {
		result[i] = ProductOutputDto{
			ID:          m[i].ID,
			UserID:      m[i].UserID,
			Name:        m[i].Name,
			Description: m[i].Description,
			File:        m[i].File,
			Version:     m[i].Version,
			Price:       services.CoinInt64ToString(m[i].Price),
		}
	}

	return result
}
