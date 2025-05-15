package service

import (
	"backend-expense-app/internals/models"
	"backend-expense-app/internals/repository"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user *models.User) (*models.User, error) {
	existingUser, err := s.repo.CheckUser(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	return s.repo.RegisterUser(user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
