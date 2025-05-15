package repository

import (
	"backend-expense-app/internals/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUserRepo(user *models.User) (*models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUsersRepo(name, email string) ([]*models.User, error) {
	var users []*models.User
	query := r.db.Model(&models.User{})

	if name != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("LOWER(email) LIKE LOWER(?)", "%"+email+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByIDRepo(id uuid.UUID) (*models.User, error) {
	var user *models.User
	if err := r.db.Where(&models.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUserRepo(user *models.User) (*models.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUserRepo(id uuid.UUID) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
