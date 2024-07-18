package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type TopUpUserDto struct {
	UserId string
	Tokens string
}

type ITopUpUserUseCase interface {
	Execute(dto TopUpUserDto) (*models.User, error)
}

type TopUpUserUseCase struct {
	userRepo repositories.UserRepository
}

func NewTopUpUserUseCase(userRepo repositories.UserRepository) *TopUpUserUseCase {
	return &TopUpUserUseCase{userRepo: userRepo}
}

func (tu *TopUpUserUseCase) Execute(in TopUpUserDto) (*models.User, error) {
	user, err := tu.userRepo.FindById(in.UserId)
	if err != nil {
		return nil, errors.New("user not found")
	}
	err = user.TopUpFromString(in.Tokens)
	if err != nil {
		return nil, err
	}
	err = tu.userRepo.Save(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
