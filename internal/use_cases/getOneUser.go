package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type GetOneUserDto struct {
	ID string
}

type IGetOneUserUseCase interface {
	Execute(dto GetOneUserDto) (*models.User, error)
}

type GetOneUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewGetOneUserUseCase(userRepo repositories.UserRepository) *GetOneUserUseCase {
	return &GetOneUserUseCase{userRepo: userRepo}
}

func (c *GetOneUserUseCase) Execute(in GetOneUserDto) (*models.User, error) {
	user, err := c.userRepo.FindById(in.ID)
	if err != nil {
		return nil, errors.New("the user does not exist")
	}

	return &user, nil
}
