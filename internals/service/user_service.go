package service

import (
	"backend-expense-app/internals/models"
	"backend-expense-app/internals/repository"
	"errors"

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

func (s *UserService) CreateUserService(user *models.User) (*models.User, error) {

	if len(user.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	existingUser, _ := s.repo.GetUserByEmailRepo(user.Email)

	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	user = &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	hashedPassword, err := user.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashedPassword

	return s.repo.CreateUserRepo(user)
}

func (s *UserService) UpdateUserService(user *models.User) (*models.User, error) {

	emailExists, _ := s.repo.GetUserByEmailRepo(user.Email)

	if emailExists != nil && emailExists.ID != user.ID {
		return nil, errors.New("email already exists")
	}

	return s.repo.UpdateUserRepo(user)
}

func (s *UserService) DeleteUserService(id uuid.UUID) error {
	userExists, _ := s.repo.GetUserByIDRepo(id)
	if userExists == nil {
		return errors.New("user not found")
	}

	return s.repo.DeleteUserRepo(id)
}

// change to validation user
