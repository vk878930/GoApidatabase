package usecase

import (
	"fmt"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
)

// transactionUsecase struct holds the reference to the transaction repository
type transactionUsecase struct {
	repo repository.TransactionRepository
}

// TransactionUseCase interface defines the methods for transaction-related business logic
type TransactionUseCase interface {
	CreateNewTransaction(transaction model.Transaction) (model.Transaction, error)
	GetAllTransaction() ([]model.Transaction, error)
	GetTransactionByID(id int) (model.Transaction, error)
}

// CreateNewTransaction handles the business logic for creating a new transaction
func (tu *transactionUsecase) CreateNewTransaction(transaction model.Transaction) (model.Transaction, error) {
	// Basic validation: Check if the transaction has necessary details
	if len(transaction.BillDetails) == 0 {
		return model.Transaction{}, fmt.Errorf("transaction must include at least one product in the bill details")
	}

	// Call the repository to create the transaction
	createdTransaction, err := tu.repo.CreateNewTransaction(transaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("failed to create transaction: %v", err)
	}

	return createdTransaction, nil
}

// GetAllTransaction retrieves all transactions and adds additional logic if needed
func (tu *transactionUsecase) GetAllTransaction() ([]model.Transaction, error) {
	transactions, err := tu.repo.GetAllTransaction()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transactions: %v", err)
	}

	// Optional: Add additional business logic here if needed
	if len(transactions) == 0 {
		fmt.Println("No transactions found")
	}

	return transactions, nil
}

// GetTransactionByID retrieves a transaction by its ID and handles error cases
func (tu *transactionUsecase) GetTransactionByID(id int) (model.Transaction, error) {
	if id <= 0 {
		return model.Transaction{}, fmt.Errorf("invalid transaction ID: %d", id)
	}

	transaction, err := tu.repo.GetTransactionByID(id)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("transaction with ID %d not found: %v", id, err)
	}

	return transaction, nil
}

// NewTransactionUseCase creates a new instance of the transactionUsecase
func NewTransactionUseCase(repo repository.TransactionRepository) TransactionUseCase {
	return &transactionUsecase{repo: repo}
}
