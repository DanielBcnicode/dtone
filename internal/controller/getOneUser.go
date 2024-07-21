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

// GetOneUser godoc
// @Sumary        Get a user
// @Description   Get user data
// @Tags          user
// @Accept        json
// @Produce       json
// @Param         user_id	path	string	true	"User ID"
// @Success 200 {object} controller.GetOneUserOutputDto
// @Failure       400
// @Failure       500
// @Security JWT
// @Router /api/v1/users/{user_id} [get]
func (gController *GetOneUserController) Handle(c *gin.Context) {
	userId := c.Param("user_id")

	dtoIn := use_cases.GetOneUserDto{ID: userId}
	user, err := gController.GetOneUserUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gController.modelToOutputDto(user))
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
