// internal/handler/user_handler.go
package handler

import (
	"backend-expense-app/internals/models"
	"backend-expense-app/internals/service"
	"encoding/json"
	"net/http"
	"regexp"
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

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	users, err := h.UserService.GetUsersService(name, email)
	if err != nil {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
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

	if id == "" || id == "undefined" {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: "ID is required",
			Status:  http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: "Invalid ID",
			Status:  http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := h.UserService.GetUserByIDService(newId)
	if err != nil {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

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

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: "Invalid request body",
			Status:  http.StatusBadRequest,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if userRequest.Name == "" || userRequest.Email == "" || userRequest.Password == "" {
		var errorResponse ErrorResponse = ErrorResponse{
			Message: "Name, email, and password are required",
			Status:  http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(userRequest.Email) {

		var errorResponse ErrorResponse = ErrorResponse{
			Message: "Invalid email format",
			Status:  http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := h.UserService.CreateUserService(&models.User{Name: userRequest.Name, Email: userRequest.Email, Password: userRequest.Password})

	if err != nil {
		var Status int

		switch {
		case err.Error() == "password must be at least 6 characters":
			Status = http.StatusBadRequest
		case err.Error() == "email already exists":
			Status = http.StatusConflict
		default:
			Status = http.StatusInternalServerError
		}

		var errorResponse ErrorResponse = ErrorResponse{
			Message: err.Error(),
			Status:  Status,
		}

		w.WriteHeader(Status)
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

// func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	id := mux.Vars(r)["id"]
// 	if id == "" || id == "undefined" {
// 		var errorResponse ErrorResponse = ErrorResponse{
// 			Message: "ID is required",
// 			Status:  http.StatusBadRequest,
// 		}

// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(errorResponse)
// 		return
// 	}
// }
