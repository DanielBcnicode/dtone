package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
)

type GetUserTransactionsDto struct {
	UserId string
}

type IGetUserTransactionsUseCase interface {
	Execute(dto GetUserTransactionsDto) ([]models.Transaction, error)
}

type GetUserTransactionsUseCase struct {
	transactionRepo repositories.TransactionRepository
}

func NewGetUserTransactionsUseCase(transactionRepo repositories.TransactionRepository) *GetUserTransactionsUseCase {
	return &GetUserTransactionsUseCase{transactionRepo: transactionRepo}
}

func (c *GetUserTransactionsUseCase) Execute(in GetUserTransactionsDto) ([]models.Transaction, error) {
	users, err := c.transactionRepo.FindAllForAUser(in.UserId, nil, nil)
	if err != nil {
		return nil, errors.New("not transaction exist")
	}

	return users, nil
}
