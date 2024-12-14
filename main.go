package main

import (
	"backend-expense-app/internals/config"
	"backend-expense-app/internals/router"
	"fmt"
	"net/http"
)

func main() {
	config.Init()

	db := config.InitDB()

	config.Migrate(db)

	appContainer := config.NewAppContainer(db)

	r := router.SetupRoutes(appContainer)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
