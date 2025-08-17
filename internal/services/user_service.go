package services

import (
	"fmt"
	"time"

	"todo-list-provider/internal/auth"
	"todo-list-provider/internal/models"
	"todo-list-provider/internal/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req *models.RegisterUserRequest) (*models.User, error) {
	if err := s.validateCreateUserRequest(req); err != nil {
		return nil, err
	}

	exists, err := s.userRepo.UserExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil, models.ErrUserExists
	}

	if !auth.IsValidPassword(req.Password) {
		return nil, models.ErrNotValidPassword
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	now := time.Now()
	user := &models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdUser, err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in database: %w", err)
	}

	return createdUser, nil
}

func (s *UserService) validateCreateUserRequest(req *models.RegisterUserRequest) error {
	if req.Email == "" {
		return models.ErrEmptyEmail
	}

	if req.Password == "" {
		return models.ErrEmptyPassword
	}

	if len(req.Password) < 6 {
		return models.ErrInvalidPassword
	}

	if req.Name == "" {
		return models.ErrEmptyName
	}

	// TODO: Add email format validation
	// TODO: Add password strength validation

	return nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
