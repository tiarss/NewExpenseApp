package repository

import (
	"backend-expense-app/internals/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) (*models.Category, error) {
	if err := r.db.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) GetCategoryByID(id uuid.UUID) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) GetCategories(name, categoryType string) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, name, categoryType).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *CategoryRepository) UpdateCategory(category *models.Category) (*models.Category, error) {
	if err := r.db.Save(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategory(category *models.Category) error {
	if err := r.db.Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
