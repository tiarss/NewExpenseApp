// internal/handler/user_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend-expense-app/internals/models"
	"backend-expense-app/internals/service"
	"backend-expense-app/internals/utils"
)

type DataResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token,omitempty"`
}

type AuthResponse struct {
	Message string       `json:"message"`
	Data    DataResponse `json:"data"`
	Status  int          `json:"status"`
}

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid request payload",
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := h.UserService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: http.StatusUnauthorized,
			Message:    "Unauthorized",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to generate token",
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := AuthResponse{
		Message: "Login successful",
		Data: DataResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
			Token: token,
		},
		Status: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	var registerRequest RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
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

	user := &models.User{
		Email: registerRequest.Email,
		Name:  registerRequest.Name,
	}

	hashedPassword, err := user.HashPassword(registerRequest.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to Hash Password",
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user.Password = hashedPassword

	user, err = h.UserService.RegisterUser(user)
	if err != nil {
		var statusCode int

		switch {
		case strings.Contains(err.Error(), "already exists"):
			statusCode = http.StatusConflict
		default:
			statusCode = http.StatusInternalServerError
		}

		// Create a dynamic error response
		errorResponse := struct {
			StatusCode int    `json:"status_code"`
			Message    string `json:"message"`
		}{
			StatusCode: statusCode,
			Message:    err.Error(),
		}

		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	response := AuthResponse{
		Message: "Registration successful",
		Data: DataResponse{
			ID:    user.ID.String(),
			Name:  user.Name,
			Email: user.Email,
		},
		Status: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
