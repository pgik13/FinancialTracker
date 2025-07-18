package routes

import (
	"trackytrack/handler"
	"trackytrack/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, transactionHandler *handler.TransactionHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("login", userHandler.LogIn).Methods("POST")

	r.HandleFunc("/transactions", transactionHandler.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions/{id}", transactionHandler.EditTransaction).Methods("PATCH")
	r.HandleFunc("/transactions/delete/{id}", transactionHandler.DeleteTransaction).Methods("DELETE")
	r.HandleFunc("/transactions/{id}", transactionHandler.GetTransactionByID).Methods("GET")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	return r
}
