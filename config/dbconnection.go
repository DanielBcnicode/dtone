package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (db *DBConnection) DatabaseConnection() (*gorm.DB, error) {
	connectString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db.Host, db.Port, db.User, db.Password, db.Database)
	gdb, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return gdb, nil
}
