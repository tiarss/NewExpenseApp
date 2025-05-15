package router

import (
	"backend-expense-app/internals/config"
	"backend-expense-app/internals/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(container *config.AppContainer) *mux.Router {
	r := mux.NewRouter()

	// Public routes that don't require authentication
	publicRoutes := r.PathPrefix("/api").Subrouter()
	publicRoutes.HandleFunc("/register", container.AuthHandler.RegisterHandler).Methods("POST")
	publicRoutes.HandleFunc("/login", container.AuthHandler.LoginHandler).Methods("POST")

	// Protected routes that require authentication
	protectedRoutes := r.PathPrefix("/api").Subrouter()
	protectedRoutes.Use(middleware.AuthMiddleware)
	protectedRoutes.HandleFunc("/users", container.UserHandler.GetUsersHandler).Methods("GET")
	protectedRoutes.HandleFunc("/users/{id}", container.UserHandler.GetUserHandler).Methods("GET")
	protectedRoutes.HandleFunc("/users", container.UserHandler.CreateUserHandler).Methods("POST")
	// r.HandleFunc("/api/expense", container.).Methods("GET")

	return r
}
