package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Password string
	User     string
	DbName   string
	SSLMode  string
}

func LoadPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
}

func ConnectToDB(config *PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.DbName, config.Port, config.SSLMode)

	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if error != nil {
		return nil, error
	}

	return db, nil
}
