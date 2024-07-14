package repositories

import "dtonetest/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByUsername(username string) (models.User, error)
	FindById(userId string) (models.User, error)
}
