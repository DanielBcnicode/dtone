package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Configuration struct {
	Database      *DBConnection
	ApiSecret     string
	TokenLifeSpan int
}

func GetConfiguration() (*Configuration, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	db := DBConnection{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	ls, err := strconv.Atoi(os.Getenv("TOKEN_LIFE_SPAN"))
	cnf := &Configuration{
		Database:      &db,
		ApiSecret:     os.Getenv("API_SECRET"),
		TokenLifeSpan: ls,
	}

	return cnf, nil
}
