package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOneUserOutputDto struct {
	CreateUserOutputDto
	Balance string `json:"balance"`
}

type GetOneUserController struct {
	GetOneUserUseCase use_cases.IGetOneUserUseCase
}

func NewGetOneUserController(
	GetOneUserUseCase use_cases.IGetOneUserUseCase) GetOneUserController {
	return GetOneUserController{
		GetOneUserUseCase: GetOneUserUseCase,
	}
}
func (gController *GetOneUserController) GetOneUser(c *gin.Context) {
	userId := c.Param("user_id")

	dtoIn := use_cases.GetOneUserDto{ID: userId}
	user, err := gController.GetOneUserUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": gController.modelToOutputDto(user)})
}

func (gController *GetOneUserController) modelToOutputDto(m *models.User) GetOneUserOutputDto {
	return GetOneUserOutputDto{
		CreateUserOutputDto: CreateUserOutputDto{
			ID:        m.ID,
			Email:     m.Email,
			Name:      m.Name,
			Telephone: m.Telephone,
		},
		Balance: m.GetBalanceFormatted(),
	}
}
