package controller

import (
	"dtonetest/internal/services"
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BuyProductInputDto struct {
	UserID string `json:"user_id" binding:"required"`
}

type BuyProductOutputDto struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	FromID          string    `json:"from_id"`
	ToID            string    `json:"to_id"`
	ProductID       string    `json:"product_id"`
	Price           string    `json:"price"`
	TransactionDate time.Time `json:"transaction_date"`
}

type BuyProductController struct {
	BuyProductUseCase use_cases.IBuyProductUseCase
}

func NewBuyProductController(
	CreateUseCase use_cases.IBuyProductUseCase,
) BuyProductController {
	return BuyProductController{
		BuyProductUseCase: CreateUseCase,
	}
}

// BuyProduct godoc
// @Sumary        Buy Product
// @Description   Buy a product
// @Tags          product
// @Accept        json
// @Produce       json
// @Param         product_id	path	string	true	"Product ID"
// @Param			userId	body		controller.BuyProductInputDto	true	"Buy Product"
// @Success 200 {object} controller.BuyProductOutputDto
// @Failure       400
// @Failure       500
// @Security JWT
// @Router /api/v1/products/{product_id}/buy [post]
func (cController *BuyProductController) HandleBuy(c *gin.Context) {
	var input BuyProductInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productId := c.Param("product_id")
	dtoIn := use_cases.BuyProductDto{
		UserID:    input.UserID,
		ProductID: productId,
		Type:      models.TransactionTypeBuy,
	}
	transaction, err := cController.BuyProductUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cController.modelToOutputDto(transaction))
}

// GiftProduct godoc
// @Sumary        Gift Product
// @Description   Gift a product
// @Tags          product
// @Accept        json
// @Produce       json
// @Param         product_id	path	string	true	"Product ID"
// @Param			userId	body		controller.BuyProductInputDto	true	"Buy Product"
// @Success 200 {object} controller.BuyProductOutputDto
// @Failure       400
// @Failure       500
// @Security JWT
// @Router /api/v1/products/{product_id}/gift [post]
func (cController *BuyProductController) HandleGift(c *gin.Context) {
	var input BuyProductInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	productId := c.Param("product_id")
	dtoIn := use_cases.BuyProductDto{
		UserID:    input.UserID,
		ProductID: productId,
		Type:      models.TransactionTypeGift,
	}
	transaction, err := cController.BuyProductUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, cController.modelToOutputDto(transaction))
}

func (cController *BuyProductController) modelToOutputDto(m *models.Transaction) BuyProductOutputDto {
	return BuyProductOutputDto{
		ID:              m.ID,
		Type:            m.Type,
		FromID:          m.FromID,
		ToID:            m.ToID,
		ProductID:       m.ProductID,
		Price:           services.CoinInt64ToString(m.Price),
		TransactionDate: m.TransactionDate,
	}
}
