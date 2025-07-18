package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trackytrack/middleware"
	"trackytrack/models"
	"trackytrack/services"

	"github.com/gorilla/mux"
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
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "transactionID is required", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid transactionID", 400)
	}

	transaction, err := h.Service.GetTransactionByID(uint(idInt))
	if err != nil {
		http.Error(w, "could not fetch transaction", 400)
		return
	}

	json.NewEncoder(w).Encode(transaction)
}

func (h *TransactionHandler) EditTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "transactionID required", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid transactionID", 400)
		return
	}

	transaction, err := h.Service.GetTransactionByID(uint(idInt))
	if err != nil {
		http.Error(w, "transaction not found", 404)
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unathorised", 401)
		return
	}

	if transaction.UserID != userID {
		http.Error(w, "you do not have permission to edit this transaction", 403)
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "invalid request bodt", 400)
		return
	}

	err = h.Service.EditTransaction(uint(idInt), updates)
	if err != nil {
		http.Error(w, "could not update transaction", 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Transaction Updated"))
}

func (h *TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transaction

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "transactionID is required", 400)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "incorrect transactionID", 400)
		return
	}

	transactions, err := h.Service.GetTransactionByID(uint(idInt))
	if err != nil {
		http.Error(w, "transaction not found", 404)
		return
	}

	userID, err := middleware.GetUserIDFromToken(r)
	if err != nil {
		http.Error(w, "unauthorised", http.StatusUnauthorized)
		return
	}

	if transactions.UserID != userID {
		http.Error(w, "You do not have permission to delete this transaction", http.StatusForbidden)
		return
	}

	err = h.Service.DeleteTransaction(&transaction, uint(idInt))
	if err != nil {
		http.Error(w, "unable to delete transaction", 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("transaction deleted"))
}
