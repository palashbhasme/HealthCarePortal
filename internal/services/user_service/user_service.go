package user_service

import (
	"fmt"
	"strconv"

	"github.com/palashbhasme/healthcare-portal/internal/api/dto/mapper"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/request"
	"github.com/palashbhasme/healthcare-portal/internal/api/dto/response"
	"github.com/palashbhasme/healthcare-portal/internal/domain/repository"
	"github.com/palashbhasme/healthcare-portal/utils"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(userRequest *request.UserRequest) (*response.UserResponse, error) {
	user := mapper.UserToModel(userRequest)
	exists, err := s.userRepo.CheckUserExists(user.Username)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, fmt.Errorf("user already exists")
	}
	hashed_password, err := utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashed_password

	newUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	userResponse := mapper.UserToResponse(newUser)
	return userResponse, nil
}

func (s *UserService) GetUserByID(idStr string) (*response.UserResponse, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	user, err := s.userRepo.GetUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	userResponse := mapper.UserToResponse(user)
	return userResponse, nil
}

func (s *UserService) UpdateUserById(idStr string, updates map[string]interface{}) (*response.UserResponse, error) {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID")
	}

	user, err := s.userRepo.UpdateUserById(uint(id), updates)
	if err != nil {
		return nil, err
	}
	userResponse := mapper.UserToResponse(user)
	return userResponse, nil
}

func (s *UserService) DeleteUserById(idStr string) error {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user ID")
	}

	err = s.userRepo.DeleteUserById(uint(id))
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) LoginUser(userRequest *request.UserLoginRequest) (*response.UserResponse, error) {
	user, err := s.userRepo.GetUserByName(userRequest.Username)
	if err != nil {
		return nil, err
	}
	err = utils.ComparePasswordHash(userRequest.Password, user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	userResponse := mapper.UserToResponse(user)
	return userResponse, nil
}
