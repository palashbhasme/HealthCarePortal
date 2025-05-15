package repository

import (
	"github.com/palashbhasme/healthcare-portal/internal/domain/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	result := r.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) UpdateUserById(id uint, updates map[string]interface{}) (*models.User, error) {
	var user models.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) DeleteUserById(id uint) error {
	var user models.User

	err := r.db.Delete(&user, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByName(username string) (*models.User, error) {
	var user models.User

	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CheckUserExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)`
	err := r.db.Raw(query, username).Scan(&exists).Error
	if err != nil {
		return false, err
	}
	return exists, nil
}
