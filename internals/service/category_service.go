package service

import "backend-expense-app/internals/repository"

type CategoryService struct {
	repo *repository.CategoryRepository
}

type SubCategoryService struct {
	repo *repository.SubCategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}
