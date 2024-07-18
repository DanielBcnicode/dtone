package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TopUpUserInputDto struct {
	UserID string `json:"user_id"`
	Amount string `json:"amount" binding:"required,gt=0"`
}

type TopUpUserOutputDto struct {
	CreateUserOutputDto
	Balance string `json:"balance"`
}

type TopUpUserController struct {
	TopUpUserUseCase use_cases.ITopUpUserUseCase
}

func NewTopUpUserController(
	topUpUseCase use_cases.ITopUpUserUseCase,
) TopUpUserController {
	return TopUpUserController{
		TopUpUserUseCase: topUpUseCase,
	}
}

func (tController *TopUpUserController) Handle(c *gin.Context) {
	var input TopUpUserInputDto
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoIn := use_cases.TopUpUserDto{
		UserId: c.Param("user_id"),
		Tokens: input.Amount,
	}

	user, err := tController.TopUpUserUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"user": tController.modelToOutputDto(user)})
}

func (tController *TopUpUserController) modelToOutputDto(m *models.User) TopUpUserOutputDto {
	return TopUpUserOutputDto{
		CreateUserOutputDto: CreateUserOutputDto{
			ID:        m.ID,
			Email:     m.Email,
			Name:      m.Name,
			Telephone: m.Telephone,
		},
		Balance: m.GetBalanceFormatted(),
	}
}
