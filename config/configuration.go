package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strconv"
)

type OTLConfig struct {
	ServiceName       string
	CollectorUrl      string
	InsecureCollector string
}
type Configuration struct {
	Database         *DBConnection
	ApiSecret        string
	TokenLifeSpan    int
	OTL              OTLConfig
	FolderRepository string
}

func (c *Configuration) checkFolderRepository() error {
	path := filepath.Join(".", c.FolderRepository)
	fi, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("it can not create folder %s", path)
		}
		fi, err = os.Stat(path)
	}
	if err != nil || !fi.IsDir() {
		return fmt.Errorf("%s is not a folder", path)
	}

	return nil
}

func GetConfiguration() (*Configuration, error) {
	err := godotenv.Load()
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
		Database:         &db,
		ApiSecret:        os.Getenv("API_SECRET"),
		TokenLifeSpan:    ls,
		FolderRepository: os.Getenv("FOLDER_REPOSITORY"),
		OTL: OTLConfig{
			ServiceName:       os.Getenv("SERVICE_NAME"),
			CollectorUrl:      os.Getenv("COLLECTOR_URL"),
			InsecureCollector: os.Getenv("INSECURE_COLLECTOR"),
		},
	}

	err = cnf.checkFolderRepository()
	if err != nil {
		return nil, err
	}

	return cnf, nil
}
