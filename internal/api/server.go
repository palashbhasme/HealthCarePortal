package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/palashbhasme/healthcare-portal/config"
	"github.com/palashbhasme/healthcare-portal/internal/api/handlers"
	"github.com/palashbhasme/healthcare-portal/internal/domain/repository"
	"github.com/palashbhasme/healthcare-portal/internal/services/patient_service"
	"github.com/palashbhasme/healthcare-portal/internal/services/user_service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Server(logger *zap.Logger, db *gorm.DB) error {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	patientRepo := repository.NewPatientRepository(db)

	userService := user_service.NewUserService(userRepo)
	patientService := patient_service.NewPatientService(patientRepo)
	authConfig := config.NewAuthConfig(os.Getenv("JWT_SECRET"))

	handlers.NewHandler(router, logger, userService, patientService, authConfig)

	if err := router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
