package config

import (
	"os"
)

type Config struct {
	CloudStorageBucket string
	DBUsername         string
	DBPassword         string
	DBHost             string
	DBPort             string
	DBName             string
}

func LoadConfig() *Config {
	return &Config{
		CloudStorageBucket: os.Getenv("CLOUD_STORAGE_BUCKET"),
		DBUsername:         os.Getenv("DB_USERNAME"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBName:             os.Getenv("DB_NAME"),
	}
}
