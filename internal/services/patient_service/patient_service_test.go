// services/patient_service/patient_service_test.go
package patient_service

import (
	"errors"
	"testing"
	"time"

	"github.com/palashbhasme/healthcare-portal/internal/api/dto/request"
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"github.com/palashbhasme/healthcare-portal/internal/services/patient_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePatient_Success(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	dob, err := time.Parse("2006-01-02", "1985-07-10")
	assert.NoError(t, err)

	email := "patient@example.com"

	patientReq := &request.PatientRequest{
		FirstName:      "Patient1",
		LastName:       "Louis",
		DOB:            "1985-07-10",
		Email:          email,
		Gender:         "male",
		PhoneNumber:    "1234567890",
		Address:        "Church Street",
		MedicalHistory: "Diabetes",
	}

	mockPatientModel := &models.Patient{
		FirstName:      "Patient1",
		LastName:       "Louis",
		DOB:            dob,
		Email:          &email,
		Gender:         "male",
		PhoneNumber:    "1234567890",
		Address:        "Church Street",
		MedicalHistory: "Diabetes",
	}

	mockRepo.On("CreatePatient", mock.AnythingOfType("*models.Patient")).Return(mockPatientModel, nil)

	result, err := service.CreatePatient(patientReq, "receptionist")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Patient1", result.FirstName)
	assert.Equal(t, &email, result.Email)
	mockRepo.AssertExpectations(t)
}

func TestCreatePatient_PermissionDenied(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)
	email := "patient@example.com"

	patientReq := &request.PatientRequest{
		FirstName:      "Patient1",
		LastName:       "Louis",
		DOB:            "1985-07-10",
		Email:          email,
		Gender:         "male",
		PhoneNumber:    "1234567890",
		Address:        "Church Street",
		MedicalHistory: "Diabetes",
	}

	result, err := service.CreatePatient(patientReq, "doctor") // Doctor cannot create patients

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "permission denied", err.Error())
}

func TestGetPatientById_Success(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	mockPatientModel := &models.Patient{
		FirstName: "Jane",
		LastName:  "Doe",
	}

	mockRepo.On("GetPatientById", uint(1)).Return(mockPatientModel, nil)

	result, err := service.GetPatientById("1", "receptionist")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Jane", result.FirstName)
	assert.Equal(t, "Doe", result.LastName)
	mockRepo.AssertExpectations(t)
}

func TestGetPatientById_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	mockRepo.On("GetPatientById", uint(99)).Return(nil, errors.New("not found"))

	result, err := service.GetPatientById("99", "receptionist")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "not found", err.Error())
}

func TestUpdatePatientById_Success(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	updates := map[string]interface{}{"FirstName": "John"}
	mockPatientModel := &models.Patient{
		FirstName: "John",
		LastName:  "Doe",
	}

	mockRepo.On("UpdatePatientById", uint(1), updates).Return(mockPatientModel, nil)

	result, err := service.UpdatePatientById("1", updates, "receptionist")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John", result.FirstName)
	mockRepo.AssertExpectations(t)
}

// To simulate a successful update using doctor role
func TestDoctorUpdatePatientById_Success(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	updates := map[string]interface{}{"FirstName": "John"}
	mockPatientModel := &models.Patient{
		FirstName: "John",
		LastName:  "Doe",
	}

	mockRepo.On("UpdatePatientById", uint(1), updates).Return(mockPatientModel, nil)

	result, err := service.UpdatePatientById("1", updates, "doctor")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "John", result.FirstName)
	mockRepo.AssertExpectations(t)
}

func TestUpdatePatientById_NotFound(t *testing.T) {
	mockRepo := new(mocks.MockPatientRepository)
	service := NewPatientService(mockRepo)

	updates := map[string]interface{}{"FirstName": "John"}

	mockRepo.On("UpdatePatientById", uint(99), updates).Return(nil, errors.New("not found"))

	result, err := service.UpdatePatientById("99", updates, "receptionist")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "not found", err.Error())
}
