package repo

import (
	"trackytrack/database"
	"trackytrack/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	EditTransaction(transaction uint, updates map[string]interface{}) error
	DeleteTransaction(id uint) error
	DuplicateCheck(transaction *models.Transaction) bool
	GetTransactionByID(id uint) (models.Transaction, error)
}

type TransactionRepo struct {
}

func (r *TransactionRepo) CreateTransaction(transaction *models.Transaction) error {
	return database.DB.Create(transaction).Error
}

func (r *TransactionRepo) EditTransaction(transactionID uint, updates map[string]interface{}) error {
	err := database.DB.Save(transactionID).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepo) DeleteTransaction(transaction *models.Transaction, id uint) error {
	err := database.DB.Delete(&models.Transaction{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepo) GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.Preload("User").Where("id = ?", id).Find(&transaction).Error
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
