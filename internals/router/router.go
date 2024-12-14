package router

import (
	"backend-expense-app/internals/config"
	// "backend-expense-app/internals/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes(container *config.AppContainer) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/register", container.UserHandler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/api/login", container.UserHandler.LoginUserHandler).Methods("POST")

	// r.Use(middleware.AuthMiddleware)

	return r
}
