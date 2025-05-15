package config

import (
	"backend-expense-app/internals/handler"
	"backend-expense-app/internals/repository"
	"backend-expense-app/internals/service"

	"gorm.io/gorm"
)

type AppContainer struct {
	UserHandler     *handler.UserHandler
	AuthHandler     *handler.AuthHandler
	CategoryHandler *handler.CategoryHandler
}

func NewAppContainer(db *gorm.DB) *AppContainer {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(*userService)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(*authService)
	// Add other handlers here
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(*categoryService)

	return &AppContainer{
		UserHandler:     userHandler,
		AuthHandler:     authHandler,
		CategoryHandler: categoryHandler,
	}
}
