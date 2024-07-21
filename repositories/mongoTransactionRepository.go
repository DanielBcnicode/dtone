package repositories

import (
	"dtonetest/models"
	"gorm.io/gorm"
	"time"
)

type MongoTransactionRepository struct {
	Database *gorm.DB
}

func NewMongoTransactionRepository(database *gorm.DB) *MongoTransactionRepository {
	return &MongoTransactionRepository{Database: database}
}

func (tr *MongoTransactionRepository) Save(transaction *models.Transaction) error {
	tx := tr.Database.Save(transaction)
	return tx.Error
}
func (tr *MongoTransactionRepository) FindAllByFromID(fromID string, fromDate *time.Time, toDate *time.Time, limit int, page int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var tx *gorm.DB
	if fromDate != nil && toDate != nil {
		tx = tr.Database.
			Scopes(newPaginate(limit, page).getOrmPaginate).
			Find(&transactions, "from_id = ?", fromID).Where("transaction_date BETWEEN ? AND ?", fromDate, toDate)
	} else {
		tx = tr.Database.
			Scopes(newPaginate(limit, page).getOrmPaginate).
			Find(&transactions, "from_id = ?", fromID)
	}

	return transactions, tx.Error
}
func (tr *MongoTransactionRepository) FindAllByToID(toID string, fromDate *time.Time, toDate *time.Time, limit int, page int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var tx *gorm.DB
	if fromDate != nil && toDate != nil {
		tx = tr.Database.
			Scopes(newPaginate(limit, page).getOrmPaginate).
			Find(&transactions, "to_id = ?", toID).Where("transaction_date BETWEEN ? AND ?", fromDate, toDate)
	} else {
		tx = tr.Database.
			Scopes(newPaginate(limit, page).getOrmPaginate).
			Find(&transactions, "to_id = ?", toID)
	}

	return transactions, tx.Error
}

func (tr *MongoTransactionRepository) FindAllForAUser(userID string, fromDate *time.Time, toDate *time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var tx *gorm.DB
	if fromDate != nil && toDate != nil {
		tx = tr.Database.
			Find(&transactions, "to_id = ? OR from_id = ?", userID, userID).Where("transaction_date BETWEEN ? AND ?", fromDate, toDate)
	} else {
		tx = tr.Database.
			Find(&transactions, "to_id = ? OR from_id = ?", userID, userID)
	}

	return transactions, tx.Error
}

func (tr *MongoTransactionRepository) FindOneTransaction(productId, fromID, toID string) (*models.Transaction, error) {
	var transaction models.Transaction
	tx := tr.Database.
		First(&transaction, "product_id = ? AND from_id = ? AND to_id = ?", productId, fromID, toID)

	return &transaction, tx.Error
}
