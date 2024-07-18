package main

import (
	"database/sql"
	"dtonetest/config"
	"dtonetest/models"
	"log"
)

func main() {
	cnf, err := config.GetConfiguration()
	if err != nil {
		panic(err)
	}

	db, err := cnf.Database.DatabaseConnection()
	if err != nil {
		panic(err)
	}
	connection, err := db.DB()
	defer func(connection *sql.DB) {
		err := connection.Close()
		if err != nil {
			log.Print(err)
		}
	}(connection)

	err = db.AutoMigrate(models.User{}, models.Product{})
	if err != nil {
		panic(err)
	}
}
