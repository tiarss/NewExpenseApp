package repository

import (
	"backend-expense-app/internals/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategoryRepository struct {
	db *gorm.DB
}

func NewSubCategoryRepository(db *gorm.DB) *SubCategoryRepository {
	return &SubCategoryRepository{db: db}
}

func (r *SubCategoryRepository) CreateSubCategory(subCategory *models.SubCategory) (*models.SubCategory, error) {
	if err := r.db.Create(&subCategory).Error; err != nil {
		return nil, err
	}

	return subCategory, nil
}

// TODO: fix this , this is wrong
func (r *SubCategoryRepository) GetSubCategoryByID(id uuid.UUID) (*models.SubCategory, error) {
	var subCategory models.SubCategory

	if err := r.db.First(&subCategory, id).Error; err != nil {
		return nil, err
	}

	return &subCategory, nil
}
