package use_cases

import (
	"dtonetest/internal/services"
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
	"time"
)

type TopUpUserDto struct {
	UserId string
	Tokens string
}

type ITopUpUserUseCase interface {
	Execute(dto TopUpUserDto) (*models.User, error)
}

type TopUpUserUseCase struct {
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
}

func NewTopUpUserUseCase(userRepo repositories.UserRepository, transactionRepo repositories.TransactionRepository) *TopUpUserUseCase {
	return &TopUpUserUseCase{userRepo: userRepo, transactionRepo: transactionRepo}
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

	amount, _ := services.CoinStringToInt64(in.Tokens)
	transaction := models.Transaction{
		Type:            models.TransactionTypeTopUp,
		FromID:          user.ID,
		ToID:            "",
		ProductID:       "",
		Price:           amount,
		TransactionDate: time.Time{},
	}

	err = tu.transactionRepo.Save(&transaction)
	if err != nil {
		return &user, errors.New("transaction not stored but top up is done")
	}

	return &user, nil
}
