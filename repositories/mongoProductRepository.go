package repositories

import (
	"dtonetest/models"
	"gorm.io/gorm"
)

type paginate struct {
	limit int
	page  int
}

func newPaginate(limit, page int) *paginate {
	return &paginate{limit: limit, page: page}
}
func (p *paginate) getOrmPaginate(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit
	return db.Limit(p.limit).Offset(offset)
}

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
	tx := mr.Database.First(&product, "id = ?", productID)

	return product, tx.Error
}

func (mr *MongoProductRepository) FindAll(UserId string, limit int, page int) ([]models.Product, error) {
	var products []models.Product
	var tx *gorm.DB
	if page < 0 {
		page = 0
	}
	if UserId == "" {
		tx = mr.Database.Scopes(newPaginate(limit, page).getOrmPaginate).Find(&products)
	} else {
		tx = mr.Database.Scopes(newPaginate(limit, page).getOrmPaginate).Find(&products, "user_id = ?", UserId)
	}

	return products, tx.Error
}
