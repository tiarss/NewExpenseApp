package config

import (
	"backend-expense-app/internals/handler"
	"backend-expense-app/internals/repository"
	"backend-expense-app/internals/service"

	"gorm.io/gorm"
)

type AppContainer struct {
	UserHandler *handler.UserHandler
}

func NewAppContainer(db *gorm.DB) *AppContainer {
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(*userService)

	return &AppContainer{
		UserHandler: userHandler,
	}
}
