package repositories

import (
	"dtonetest/models"
	"gorm.io/gorm"
)

type MongoProductRepository struct {
	Database *gorm.DB
}

func NewMongoProductRepository(database *gorm.DB) *MongoProductRepository {
	return &MongoProductRepository{Database: database}
}

func (mr *MongoProductRepository) Save(product *models.Product) error {
	tx := mr.Database.Save(product)
	return tx.Error
}

func (mr *MongoProductRepository) FindById(productID string) (models.Product, error) {
	product := models.Product{}
	tx := mr.Database.Find(&product, "id = ?", productID)

	return product, tx.Error
}
