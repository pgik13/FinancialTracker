package main

import (
	"fmt"
	"net/http"
	"os"
	"trackytrack/config"
	"trackytrack/database"
	"trackytrack/handler"
	"trackytrack/repo"
	"trackytrack/routes"
	"trackytrack/services"
)

func main() {

	config.LoadEnv()

	database.ConnectDB()

	userRepo := &repo.UserRepo{}

	userService := &services.UserServices{Repo: *userRepo}

	userHandler := &handler.UserHandler{Service: userService}

	router := routes.SetupRouter(userHandler)

	fmt.Println("Server running on " + os.Getenv("DB_HOST"))
	http.ListenAndServe(":3040", router)

}
