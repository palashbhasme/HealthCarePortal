package patient_service

import (
	"errors"
	"strconv"

	"github.com/palashbhasme/healthcare-portal/internal/api/dto/mapper"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/request"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/response"
	"github.com/palashbhasme/healthcare-portal/internal/domain/repository"
	"github.com/palashbhasme/healthcare-portal/internal/services"
)

type PatientService struct {
	patientRepo repository.PatientRepository
}

func NewPatientService(patientRepo repository.PatientRepository) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
	}
}

func (s *PatientService) CreatePatient(patientRequest *request.PatientRequest, role any) (*response.PatientResponse, error) {
	roleValue, ok := role.(string)
	if !ok {
		return nil, errors.New("invalid role type")
	}

	if err := services.CheckPermission(roleValue, "create_patient"); err != nil {
		return nil, errors.New("permission denied")
	}
	patient, err := mapper.PatientToModel(patientRequest)
	if err != nil {
		return nil, err
	}
	newPatient, err := s.patientRepo.CreatePatient(patient)
	if err != nil {
		return nil, err
	}
	patientResponse := mapper.PatientToResponse(newPatient)
	return patientResponse, nil
}

func (s *PatientService) UpdatePatientById(idStr string, updates map[string]interface{}, role any) (*response.PatientResponse, error) {
	roleValue, ok := role.(string)
	if !ok {
		return nil, errors.New("invalid role type")
	}

	if err := services.CheckPermission(roleValue, "update_patient"); err != nil {
		return nil, errors.New("permission denied")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid patient ID")
	}

	patient, err := s.patientRepo.UpdatePatientById(uint(id), updates)
	if err != nil {
		return nil, err
	}
	patientResponse := mapper.PatientToResponse(patient)
	return patientResponse, nil
}

func (s *PatientService) GetPatientById(idStr string, role any) (*response.PatientResponse, error) {
	roleValue, ok := role.(string)
	if !ok {
		return nil, errors.New("invalid role type")
	}

	if err := services.CheckPermission(roleValue, "view_patient"); err != nil {
		return nil, errors.New("permission denied")
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, errors.New("invalid patient ID")
	}

	patient, err := s.patientRepo.GetPatientById(uint(id))
	if err != nil {
		return nil, err
	}
	patientResponse := mapper.PatientToResponse(patient)
	return patientResponse, nil
}
