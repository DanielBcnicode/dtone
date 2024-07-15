package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserDto struct {
	Email     string
	Password  string
	Name      string
	Telephone string
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
	_, err := c.userRepo.FindByEmail(in.Email)
	if err == nil {
		return nil, errors.New("the user already exists")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user := models.User{
		Email:     in.Email,
		Password:  string(passwordHash),
		Name:      in.Name,
		Telephone: in.Telephone,
	}
	err = c.userRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
