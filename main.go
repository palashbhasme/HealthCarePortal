package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/palashbhasme/healthcare-portal/config"
	"github.com/palashbhasme/healthcare-portal/internal/api"
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"github.com/palashbhasme/healthcare-portal/utils"
	"go.uber.org/zap"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error occured while loading environment variables")
	}

	logger, err := utils.InitializeLogger()
	if err != nil {
		log.Fatal("Error occured while initializing logger")
	}

	defer logger.Sync()
	logger.Info("Logger initialized")

	postgresConfig := config.LoadPostgresConfig()

	db, err := config.ConnectToDB(postgresConfig)

	if err != nil {
		log.Fatal("Error connecting to database", zap.Error(err))
	}

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatal("Error migrating databse", zap.Error(err))
	}

	api.Server(logger, db)
	if err != nil {
		log.Fatal("Error starting server", zap.Error(err))
	}

}
