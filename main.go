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
	transactionRepo := &repo.TransactionRepo{}

	userService := &services.UserServices{Repo: *userRepo}
	transactionService := &services.TransactionService{Repo: *transactionRepo}

	userHandler := &handler.UserHandler{Service: userService}
	transactionHandler := &handler.TransactionHandler{Service: transactionService}

	router := routes.SetupRouter(userHandler, transactionHandler)

	fmt.Println("Server running on " + os.Getenv("DB_HOST"))
	http.ListenAndServe(":3040", router)

}
