package repositories

import "dtonetest/models"

type ProductRepository interface {
	Save(user *models.Product) error
	FindById(productID string) (models.Product, error)
}
