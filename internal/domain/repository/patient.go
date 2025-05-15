package repository

import (
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *patientRepository {
	return &patientRepository{
		db: db,
	}
}
func (r *patientRepository) CreatePatient(patient *models.Patient) (*models.Patient, error) {
	result := r.db.Create(patient)

	if result.Error != nil {
		return nil, result.Error
	}

	return patient, nil
}

func (r *patientRepository) UpdatePatientById(id uint, updates map[string]interface{}) (*models.Patient, error) {
	var patient models.Patient

	err := r.db.Model(&patient).Clauses(clause.Returning{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) GetPatientById(id uint) (*models.Patient, error) {
	var patient models.Patient

	err := r.db.First(&patient, id).Error
	if err != nil {
		return nil, err
	}

	return &patient, nil
}

func (r *patientRepository) DeletePatientById(id uint) error {
	var patient models.Patient

	err := r.db.Delete(&patient, id).Error
	if err != nil {
		return err
	}
	return nil
}
