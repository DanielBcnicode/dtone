package use_cases

import (
	"dtonetest/models"
	"dtonetest/repositories"
	"errors"
	"time"
)

type BuyProductDto struct {
	UserID    string
	ProductID string
	Type      string
}

type IBuyProductUseCase interface {
	Execute(dto BuyProductDto) (*models.Transaction, error)
}

type BuyProductUseCase struct {
	productRepo     repositories.ProductRepository
	userRepo        repositories.UserRepository
	transactionRepo repositories.TransactionRepository
}

func NewBuyProductUseCase(
	productRepo repositories.ProductRepository,
	userRepo repositories.UserRepository,
	transactionRepo repositories.TransactionRepository,
) *BuyProductUseCase {
	return &BuyProductUseCase{
		productRepo:     productRepo,
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (c *BuyProductUseCase) Execute(in BuyProductDto) (*models.Transaction, error) {
	user, err := c.userRepo.FindById(in.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	product, err := c.productRepo.FindById(in.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	seller, err := c.userRepo.FindById(product.UserID)
	if err != nil {
		return nil, errors.New("seller not found")
	}

	if user.ID == seller.ID {
		return nil, errors.New("seller cannot buy products from himself")
	}

	_, err = c.transactionRepo.FindOneTransaction(product.ID, seller.ID, user.ID)
	if err == nil {
		return nil, errors.New("user has bought this product before")
	}

	transaction := models.Transaction{
		Type:            in.Type,
		FromID:          seller.ID,
		ToID:            user.ID,
		ProductID:       product.ID,
		Price:           product.Price,
		TransactionDate: time.Now(),
	}
	switch in.Type {
	case models.TransactionTypeBuy:
		if user.Balance < product.Price {
			return nil, errors.New("not enough balance")
		}
		seller.Balance += product.Price
		user.Balance -= product.Price
		transaction.Price = product.Price

		err = c.userRepo.Save(&user)
		if err != nil {
			return nil, err
		}
		err = c.userRepo.Save(&seller)
		if err != nil {
			return nil, err
		}

		break
	case models.TransactionTypeGift:
		transaction.Price = 0
		break
	default:
		return nil, errors.New("transaction type not supported")
	}

	err = c.transactionRepo.Save(&transaction)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
