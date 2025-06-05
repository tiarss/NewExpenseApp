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

func (r *CategoryRepository) CreateCategoryRepo(category *models.Category) (*models.Category, error) {
	if len(category.SubCategories) > 0 {
		for i := range category.SubCategories {
			if category.SubCategories[i].ID == uuid.Nil {
				category.SubCategories[i].ID = uuid.New()
			}
			category.SubCategories[i].CategoryID = category.ID
		}
	}
	tx := r.db.Begin()
	if err := tx.Create(&category).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return category, nil
}

func (r *CategoryRepository) GetCategoryByIDRepo(id uuid.UUID) (*models.Category, error) {
	var category *models.Category
	if err := r.db.Where(&models.Category{ID: id}).Preload("SubCategories").First(&category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepository) GetCategoriesRepo(name, categoryType string) ([]*models.Category, error) {
	var category []*models.Category

	query := r.db

	if name != "" {
		query = query.Where("name = ?", name)
	}

	if categoryType != "" {
		query = query.Where("category_type = ?", categoryType)
	}

	if err := query.Preload("SubCategories").Find(&category).Error; err != nil {
		return nil, err
	}

	if len(category) == 0 {
		return make([]*models.Category, 0), nil
	}

	return category, nil
}

func (r *CategoryRepository) UpdateCategoryRepo(category *models.Category) (*models.Category, error) {
	if err := r.db.Save(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *CategoryRepository) DeleteCategoryRepo(id uuid.UUID) error {
	if err := r.db.Delete(&models.Category{}, &id).Error; err != nil {
		return err
	}

	return nil
}
