package service

import (
	"backend-expense-app/internals/models"
	"backend-expense-app/internals/repository"
	"errors"

	"github.com/google/uuid"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

type SubCategoryService struct {
	repo *repository.SubCategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategoryService(category *models.Category) (*models.Category, error) {
	createdCategory, err := s.repo.CreateCategoryRepo(category)
	if err != nil {
		return nil, err
	}
	return createdCategory, nil
}

func (s *CategoryService) GetAllCategoriesService(name string, categoryType string) ([]*models.Category, error) {
	categories, err := s.repo.GetCategoriesRepo(name, categoryType)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) GetCategoryByIDService(id uuid.UUID) (*models.Category, error) {
	category, err := s.repo.GetCategoryByIDRepo(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) DeleteCategoryIDService(id uuid.UUID) error {
	categoryExist, _ := s.repo.GetCategoryByIDRepo(id)
	if categoryExist == nil {
		return errors.New("Category not found")
	}

	// TODO: Delete SubCategory by CategoryID

	// subCategoryExist, _ := s.repo.GetSubCategoryByCategoryIDRepo(id)
	// if subCategoryExist != nil {
	// 	return errors.New("Subcategory exist")
	// }

	return s.repo.DeleteCategoryRepo(id)
}
