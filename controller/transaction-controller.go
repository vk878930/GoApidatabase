package controller

import (
	"net/http"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"
	"submission-project-enigma-laundry/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TransactionController struct holds the usecase and router group reference
type TransactionController struct {
	useCase usecase.TransactionUseCase
	rg      *gin.RouterGroup
	authMiddleware middleware.AuthMiddleware
}

// Route defines the transaction-related HTTP routes
func (tc *TransactionController) Route() {
	tc.rg.POST("/transactions", tc.authMiddleware.RequireToken("admin") , tc.createNewTransaction)     // POST request to create a new transaction
	tc.rg.GET("/transactions", tc.authMiddleware.RequireToken("admin"), tc.getAllTransactions)        // GET request to retrieve all transactions
	tc.rg.GET("/transactions/:id", tc.authMiddleware.RequireToken("admin"), tc.getTransactionByID)    // GET request to retrieve a single transaction by ID
}

// createNewTransaction handles the creation of a new transaction
func (tc *TransactionController) createNewTransaction(c *gin.Context) {
	var payload model.Transaction

	// Bind the JSON input to the transaction struct
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	transaction, err := tc.useCase.CreateNewTransaction(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create transaction: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// getAllTransactions handles retrieving all transactions from the database
func (tc *TransactionController) getAllTransactions(c *gin.Context) {
	transactions, err := tc.useCase.GetAllTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve transactions"})
		return
	}

	if len(transactions) > 0 {
		c.JSON(http.StatusOK, transactions)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List of transactions is empty"})
}

// getTransactionByID retrieves a specific transaction by its ID
func (tc *TransactionController) getTransactionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // Convert ID from string to integer
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid transaction ID"})
		return
	}

	transaction, err := tc.useCase.GetTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get transaction by ID: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// NewTransactionController creates a new instance of TransactionController with the given usecase and router group
func NewTransactionController(useCase usecase.TransactionUseCase, rg *gin.RouterGroup, am middleware.AuthMiddleware) *TransactionController {
	return &TransactionController{useCase: useCase, rg: rg, authMiddleware: am}
}