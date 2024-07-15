package controller

import (
	"dtonetest/internal/use_cases"
	"dtonetest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInputDto struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Telephone string `json:"telephone" binding:"required,e164"`
}

type CreateUserOutputDto struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

type CreateUserController struct {
	CreateUserUseCase use_cases.ICreateUserUseCase
}

func NewRegisterController(
	CreateUseCase use_cases.ICreateUserUseCase) CreateUserController {
	return CreateUserController{
		CreateUserUseCase: CreateUseCase,
	}
}
func (rController *CreateUserController) Register(c *gin.Context) {
	var input CreateUserInputDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoIn := use_cases.CreateUserDto{
		Email:     input.Email,
		Password:  input.Password,
		Name:      input.Name,
		Telephone: input.Telephone,
	}

	user, err := rController.CreateUserUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": rController.modelToOutputDto(user)})
}

func (rController *CreateUserController) modelToOutputDto(m *models.User) CreateUserOutputDto {
	return CreateUserOutputDto{
		ID:        m.ID,
		Email:     m.Email,
		Name:      m.Name,
		Telephone: m.Telephone,
	}
}
