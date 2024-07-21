package controller

import (
	"dtonetest/internal/services"
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type innerTransactionsDto struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	FromID          string    `json:"from_id"`
	ToID            string    `json:"to_id"`
	ProductID       string    `json:"product_id"`
	Price           string    `json:"price"`
	TransactionDate time.Time `json:"transaction_date"`
}
type GetUserTransactionsOutputDto struct {
	UserId       string                 `json:"userId"`
	Transactions []innerTransactionsDto `json:"transactions"`
}

type GetUserTransactionsController struct {
	GetUserTransactionsUseCase use_cases.IGetUserTransactionsUseCase
}

func NewGetUserTransactionsController(
	GetUserTransactionsUseCase use_cases.IGetUserTransactionsUseCase) GetUserTransactionsController {
	return GetUserTransactionsController{
		GetUserTransactionsUseCase: GetUserTransactionsUseCase,
	}
}

func (gController *GetUserTransactionsController) Handle(c *gin.Context) {
	userId := c.Param("user_id")
	if uuid.Validate(userId) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Id is invalid."})
		return
	}

	dtoIn := use_cases.GetUserTransactionsDto{UserId: userId}
	transactions, err := gController.GetUserTransactionsUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gController.modelToOutputDto(transactions, userId))
}

func (gController *GetUserTransactionsController) modelToOutputDto(m []models.Transaction, userId string) GetUserTransactionsOutputDto {

	var innerTransactions []innerTransactionsDto
	for _, transaction := range m {
		innerTransactions = append(innerTransactions, innerTransactionsDto{
			ID:              transaction.ID,
			Type:            transaction.Type,
			FromID:          transaction.FromID,
			ToID:            transaction.ToID,
			ProductID:       transaction.ProductID,
			Price:           services.CoinInt64ToString(transaction.Price),
			TransactionDate: transaction.TransactionDate,
		})
	}
	return GetUserTransactionsOutputDto{
		UserId:       userId,
		Transactions: innerTransactions,
	}
}
