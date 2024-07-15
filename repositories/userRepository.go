package repositories

import "dtonetest/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByEmail(email string) (models.User, error)
	FindById(userId string) (models.User, error)
}
