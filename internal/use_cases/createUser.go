package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserDto struct {
	Username string
	Password string
}

type ICreateUserUseCase interface {
	Execute(dto CreateUserDto) (*models.User, error)
}

type CreateUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewCreateUserUseCase(userRepo repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{userRepo: userRepo}
}

func (c *CreateUserUseCase) Execute(in CreateUserDto) (*models.User, error) {
	_, err := c.userRepo.FindByUsername(in.Username)
	if err == nil {
		return nil, errors.New("the user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user := models.User{
		Username: in.Username,
		Password: string(passwordHash),
	}
	err = c.userRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
