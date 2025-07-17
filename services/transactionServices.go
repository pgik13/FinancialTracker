package services

import (
	"trackytrack/models"
	"trackytrack/repo"
)

type TransactionService struct {
	Repo repo.TransactionRepo
}

func (s *TransactionService) CreateTransaction(transaction *models.Transaction) error {

	err := s.Repo.CreateTransaction(transaction)
	if err != nil {
		return err
	}
	return nil
}
