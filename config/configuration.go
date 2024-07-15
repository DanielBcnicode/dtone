package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type OTLConfig struct {
	ServiceName       string
	CollectorUrl      string
	InsecureCollector string
}
type Configuration struct {
	Database      *DBConnection
	ApiSecret     string
	TokenLifeSpan int
	OTL           OTLConfig
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
		OTL: OTLConfig{
			ServiceName:       os.Getenv("SERVICE_NAME"),
			CollectorUrl:      os.Getenv("COLLECTOR_URL"),
			InsecureCollector: os.Getenv("INSECURE_COLLECTOR"),
		},
	}

	return cnf, nil
}
