package routes

import (
	"trackytrack/handler"
	"trackytrack/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("login", userHandler.LogIn).Methods("POST")

	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	return r
}
