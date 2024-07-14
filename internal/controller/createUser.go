package controller

import (
	"dtonetest/internal/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInputDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
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
		Username: input.Username,
		Password: input.Password,
	}

	user, err := rController.CreateUserUseCase.Execute(dtoIn)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}
