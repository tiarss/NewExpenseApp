package handler

import (
	"encoding/json"
	"net/http"

	"backend-expense-app/internals/models"
	"backend-expense-app/internals/service"
)

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

		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}

func (h *CategoryHandler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {

}
