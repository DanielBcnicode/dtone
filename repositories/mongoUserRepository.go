package repositories

import (
	"dtonetest/models"
	"gorm.io/gorm"
)

type MongoUserRepository struct {
	Database *gorm.DB
}

func NewMongoUserRepository(database *gorm.DB) *MongoUserRepository {
	return &MongoUserRepository{Database: database}
}

func (mr *MongoUserRepository) Save(user *models.User) error {
	tx := mr.Database.Save(user)
	return tx.Error
}

func (mr *MongoUserRepository) FindByEmail(email string) (models.User, error) {
	user := models.User{}
	tx := mr.Database.First(&user, "email = ?", email)

	return user, tx.Error
}

func (mr *MongoUserRepository) FindById(userId string) (models.User, error) {
	user := models.User{}
	tx := mr.Database.Find(&user, "id = ?", userId)

	return user, tx.Error
}
