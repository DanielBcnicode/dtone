package repositories

import (
	"dtonetest/models"
	"time"
)

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindAllByFromID(fromID string, fromDate *time.Time, toDate *time.Time, limit int, page int) ([]models.Transaction, error)
	FindAllByToID(fromID string, fromDate *time.Time, toDate *time.Time, limit int, page int) ([]models.Transaction, error)
	FindOneTransaction(productId, fromID, toID string) (*models.Transaction, error)
	FindAllForAUser(userID string, fromDate *time.Time, toDate *time.Time) ([]models.Transaction, error)
}
