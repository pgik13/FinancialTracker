package handler

import (
	"encoding/json"
	"net/http"
	"trackytrack/middleware"
	"trackytrack/models"
	"trackytrack/services"
)

type TransactionHandler struct {
	Service *services.TransactionService
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, "invalid request body", 400)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "failed to get userID", 500)
		return
	}

	transaction.UserID = userID

	err = h.Service.CreateTransaction(&transaction)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(transaction)
}
