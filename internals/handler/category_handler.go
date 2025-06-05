package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"backend-expense-app/internals/models"
	"backend-expense-app/internals/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryData struct {
	ID            uuid.UUID         `json:"id"`
	Name          string            `json:"name"`
	CategoryType  string            `json:"type"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
	SubCategories []SubCategoryData `json:"sub_categories"`
}

type SubCategoryData struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}

type CategoryResponse struct {
	Status  int          `json:"status_code"`
	Message string       `json:"message"`
	Data    CategoryData `json:"data"`
}

type CategoriesResponse struct {
	Status  int            `json:"status_code"`
	Message string         `json:"message"`
	Data    []CategoryData `json:"data"`
}

type CategoryHandler struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{CategoryService: categoryService}
}

func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	createdCategory, err := h.CategoryService.CreateCategoryService(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create category",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var subCategories []SubCategoryData
	for _, subCategory := range createdCategory.SubCategories {
		subCategoryData := SubCategoryData{
			ID:         subCategory.ID,
			Name:       subCategory.Name,
			CategoryID: subCategory.CategoryID,
			CreatedAt:  subCategory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  subCategory.UpdatedAt.Format(time.RFC3339),
		}
		subCategories = append(subCategories, subCategoryData)
	}

	createdCategoryData := CategoryData{
		ID:            createdCategory.ID,
		Name:          createdCategory.Name,
		CategoryType:  createdCategory.CategoryType,
		CreatedAt:     createdCategory.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     createdCategory.UpdatedAt.Format(time.RFC3339),
		SubCategories: subCategories,
	}

	response := CategoryResponse{
		Status:  http.StatusCreated,
		Message: "Category created successfully",
		Data:    createdCategoryData,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
	return
}

func (h *CategoryHandler) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	categoryType := r.URL.Query().Get("type")

	categories, err := h.CategoryService.GetAllCategoriesService(name, categoryType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve categories",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var categoriesData []CategoryData

	if len(categories) == 0 {
		w.WriteHeader(http.StatusOK)
		errorResponse := CategoriesResponse{
			Status:  http.StatusOK,
			Message: "Categories retrieved successfully",
			Data:    []CategoryData{},
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	for _, category := range categories {
		var subCategories []SubCategoryData
		for _, subCategory := range category.SubCategories {
			subCategoryData := SubCategoryData{
				ID:         subCategory.ID,
				Name:       subCategory.Name,
				CategoryID: subCategory.CategoryID,
				CreatedAt:  subCategory.CreatedAt.Format(time.RFC3339),
				UpdatedAt:  subCategory.UpdatedAt.Format(time.RFC3339),
			}
			subCategories = append(subCategories, subCategoryData)
		}
		categoryData := CategoryData{
			ID:            category.ID,
			Name:          category.Name,
			CategoryType:  category.CategoryType,
			CreatedAt:     category.CreatedAt.Format(time.RFC3339),
			UpdatedAt:     category.UpdatedAt.Format(time.RFC3339),
			SubCategories: subCategories,
		}
		categoriesData = append(categoriesData, categoryData)
	}

	response := CategoriesResponse{
		Status:  http.StatusOK,
		Message: "Categories retrieved successfully",
		Data:    categoriesData,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

func (h *CategoryHandler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" || id == "undefined" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "ID is required",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID format",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	category, err := h.CategoryService.GetCategoryByIDService(newId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Category no Found",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	var subCategories []SubCategoryData
	for _, subCategory := range category.SubCategories {
		subCategoryData := SubCategoryData{
			ID:         subCategory.ID,
			Name:       subCategory.Name,
			CategoryID: subCategory.CategoryID,
			CreatedAt:  subCategory.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  subCategory.UpdatedAt.Format(time.RFC3339),
		}
		subCategories = append(subCategories, subCategoryData)
	}

	categoryData := CategoryData{
		ID:            category.ID,
		Name:          category.Name,
		CategoryType:  category.CategoryType,
		CreatedAt:     category.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     category.UpdatedAt.Format(time.RFC3339),
		SubCategories: subCategories,
	}

	response := CategoryResponse{
		Status:  http.StatusOK,
		Message: "Category retrieved successfully",
		Data:    categoryData,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}

func (h *CategoryHandler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if id == "" || id == "undefined" {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "ID is required",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorResponse := ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID format",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	err = h.CategoryService.DeleteCategoryIDService(newId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Category not Found",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := CategoryResponse{
		Status:  http.StatusOK,
		Message: "Category deleted successfully",
		Data:    CategoryData{},
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	return
}
