package mocks

import (
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockPatientRepository struct {
	mock.Mock
}

func (m *MockPatientRepository) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	args := m.Called(patient)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Patient), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPatientRepository) GetPatientById(id uint) (*models.Patient, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Patient), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPatientRepository) UpdatePatientById(id uint, updates map[string]interface{}) (*models.Patient, error) {
	args := m.Called(id, updates)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Patient), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockPatientRepository) DeletePatientById(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
