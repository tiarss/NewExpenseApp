package service

import (
	"backend-expense-app/internals/models"
	"backend-expense-app/internals/repository"

	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUsersService(name string, email string) ([]*models.User, error) {
	users, err := s.repo.GetUsersRepo(name, email)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByIDService(id uuid.UUID) (*models.User, error) {
	user, err := s.repo.GetUserByIDRepo(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// change to validation user
