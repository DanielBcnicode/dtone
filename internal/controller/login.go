package controller

import (
	"dtonetest/internal/use_cases"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInputDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginController struct {
	LoginUseCase use_cases.ILoginUseCase
}

func NewLoginController(useCase use_cases.ILoginUseCase) LoginController {
	return LoginController{LoginUseCase: useCase}
}

func (lc LoginController) Login(c *gin.Context) {
	var input LoginInputDto
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dtoIn := use_cases.LoginDto{
		Email:    input.Email,
		Password: input.Password,
	}

	response, err := lc.LoginUseCase.Execute(dtoIn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": response.Jwt})
}
