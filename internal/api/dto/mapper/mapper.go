package mapper

import (
	"time"

	"github.com/palashbhasme/healthcare-portal/internal/api/dto/request"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/response"
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
)

func UserToResponse(user *models.User) *response.UserResponse {
	return &response.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role)}
}

func PatientToResponse(patient *models.Patient) *response.PatientResponse {
	return &response.PatientResponse{
		ID:             patient.ID,
		FirstName:      patient.FirstName,
		LastName:       patient.LastName,
		Email:          patient.Email,
		PhoneNumber:    patient.PhoneNumber,
		Address:        patient.Address,
		MedicalHistory: patient.MedicalHistory,
		DOB:            patient.DOB,
	}
}

func UserToModel(userRequest *request.UserRequest) *models.User {
	return &models.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
		Role:     models.Role(userRequest.Role),
	}
}

func PatientToModel(patientRequest *request.PatientRequest) (*models.Patient, error) {
	//parsing date from string to time.Time
	dob, err := time.Parse("2006-01-02", patientRequest.DOB)
	if err != nil {
		return nil, err
	}

	//email is nil pointer if no email is provided
	var email *string
	if patientRequest.Email != "" {
		email = &patientRequest.Email
	}
	return &models.Patient{
		FirstName:      patientRequest.FirstName,
		LastName:       patientRequest.LastName,
		DOB:            dob,
		Email:          email,
		Gender:         patientRequest.Gender,
		PhoneNumber:    patientRequest.PhoneNumber,
		Address:        patientRequest.Address,
		MedicalHistory: patientRequest.MedicalHistory,
	}, nil
}
