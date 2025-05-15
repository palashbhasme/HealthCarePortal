package repository

import "github.com/palashbhasme/healthcare-portal/internal/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUserById(id uint, updates map[string]interface{}) (*models.User, error)
	DeleteUserById(id uint) error
	GetUserByName(username string) (*models.User, error)
	CheckUserExists(username string) (bool, error)
}

type PatientRepository interface {
	CreatePatient(patient *models.Patient) (*models.Patient, error)
	GetPatientById(id uint) (*models.Patient, error)
	UpdatePatientById(id uint, updates map[string]interface{}) (*models.Patient, error)
	DeletePatientById(id uint) error
}
