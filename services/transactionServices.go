package services

import (
	"errors"
	"strings"
	"trackytrack/models"
	"trackytrack/repo"
)

type TransactionService struct {
	Repo repo.TransactionRepo
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {

	if strings.TrimSpace(transaction.Type) == "" {
		return errors.New("type cannot be empty")
	}

	if strings.TrimSpace(transaction.Category) == "" {
		return errors.New("category cannot be empty")
	}

	if transaction.Amount == 0 {
		return errors.New("amount cannot be empty")
	}

	err := s.Repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) GetTransactionByID(id uint) (*models.Transaction, error) {
	return s.Repo.GetTransactionByID(id)
}

func (s *TransactionService) EditTransaction(transactionID uint, updates map[string]interface{}) error {
	err := s.Repo.EditTransaction(transactionID, updates)
	if err != nil {
		return err
	}
	return nil
}

func (s *TransactionService) DeleteTransaction(transaction *models.Transaction, transactionID uint) error {
	err := s.Repo.DeleteTransaction(transaction, transactionID)
	if err != nil {
		return err
	}

	return nil
}
