// internal/handler/user_handler.go
package handler

import (
	"backend-expense-app/internals/service"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersResponse struct {
	Message string     `json:"message"`
	Data    []UserData `json:"data"`
	Status  int        `json:"status"`
}

type UserResponse struct {
	Message string   `json:"message"`
	Data    UserData `json:"data"`
	Status  int      `json:"status"`
}

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	users, err := h.UserService.GetUsersService(name, email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var Data []UserData

	for _, user := range users {
		userData := UserData{
			ID:        user.ID.String(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		}
		Data = append(Data, userData)
	}

	usersResponse := UsersResponse{
		Message: "Success",
		Data:    Data,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersResponse)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	fmt.Println(id)
	if id == "" {
		var errorResponse struct {
			StatusCode int    `json:"status"`
			Message    string `json:"message"`
		}
		errorResponse.StatusCode = http.StatusBadRequest
		errorResponse.Message = "ID is required"
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		var errorResponse struct {
			StatusCode int    `json:"status"`
			Message    string `json:"message"`
		}
		errorResponse.StatusCode = http.StatusBadRequest
		errorResponse.Message = "Invalid UUID format"
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := h.UserService.GetUserByIDService(newId)
	if err != nil {
		var errorResponse struct {
			StatusCode int    `json:"status"`
			Message    string `json:"message"`
		}
		errorResponse.StatusCode = http.StatusInternalServerError
		errorResponse.Message = "No User Found"
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	userData := UserData{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}

	usersResponse := UserResponse{
		Message: "Success",
		Data:    userData,
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usersResponse)
}
