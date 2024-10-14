package repository

import (
	"database/sql"
	"submission-project-enigma-laundry/model"

	"fmt"
)

type transactionRepository struct {
	db                 *sql.DB
	customerRepository CustomerRepository
	employeeRepository EmployeeRepository
}

type TransactionRepository interface {
	CreateNewTransaction(transaction model.Transaction) (model.Transaction, error)
	GetAllTransaction() ([]model.Transaction, error)
	GetTransactionByID(id int) (model.Transaction, error)
}

func (trs *transactionRepository) CreateNewTransaction(transaction model.Transaction) (model.Transaction, error) {

	var transactionID int

	tx, err := trs.db.Begin()
	if err != nil {
		return model.Transaction{}, err
	}

	err = tx.QueryRow(`
		INSERT INTO transaction (bill_date, entry_date, finish_date, employee_id, customer_id)
		VALUES ($1, $2, $3, $4, $5) RETURNING transaction_id;`,
		transaction.BillDate, transaction.EntryDate, transaction.FinishDate, transaction.EmployeeID, transaction.CustomerID).Scan(&transactionID)

	if err != nil {
		tx.Rollback()
		fmt.Printf("Error inserting into transaction: %v\n", err)
		return model.Transaction{}, err
	}

	transaction.Transaction_id = transactionID

	// Loop through BillDetails and insert into bill_details table
	for _, detail := range transaction.BillDetails {
		var productPrice int

		// Fetch product price from the product table based on product_id
		err := tx.QueryRow(`
			SELECT price FROM product WHERE product_id = $1;`,
			detail.ProductID).Scan(&productPrice)

		if err != nil {
			tx.Rollback()
			fmt.Printf("Error fetching product price: %v\n", err) // Log the error
			return model.Transaction{}, err
		}

		// Insert into the bill_details table
		_, err = tx.Exec(`
			INSERT INTO bill_details (transaction_id, product_id, product_price, qty)
			VALUES ($1, $2, $3, $4);`,
			transactionID, detail.ProductID, productPrice, detail.Qty)

		if err != nil {
			tx.Rollback()
			fmt.Printf("Error inserting into bill_details: %v\n", err) // Log the error
			return model.Transaction{}, err
		}
	}
	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error committing transaction: %v\n", err) // Log the error
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (trs *transactionRepository) GetAllTransaction() ([]model.Transaction, error) {

	var transactions []model.Transaction

	rows, err := trs.db.Query(`SELECT 
			t.transaction_id, 
			t.bill_date, 
			t.entry_date, 
			t.finish_date, 
			t.employee_id, 
			t.customer_id,
			COALESCE(SUM(b.product_price * b.qty), 0) AS total_bill
		FROM transaction t
		LEFT JOIN bill_details b ON t.transaction_id = b.transaction_id
		GROUP BY t.transaction_id;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction

		err := rows.Scan(&transaction.Transaction_id, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.EmployeeID, &transaction.CustomerID, &transaction.TotalBill)

		if err != nil {
			return nil, err
		}

		// Fetch Customer details using customerRepository
		customer, err := trs.customerRepository.GetCustByID(transaction.CustomerID)
		if err != nil {
			return nil, err
		}
		transaction.Customer = customer

		// Fetch Customer details using customerRepository
		employee, err := trs.employeeRepository.GetEmployeeByID(transaction.EmployeeID)
		if err != nil {
			return nil, err
		}
		transaction.Employee = employee

		billRows, err := trs.db.Query(`SELECT id, transaction_id, product_id, product_price, qty FROM bill_details WHERE transaction_id=$1;`, transaction.Transaction_id)

		if err != nil {
			return nil, err
		}

		for billRows.Next() {
			var detail model.BillDetails
			err := billRows.Scan(&detail.ID, &detail.BillID, &detail.ProductID, &detail.ProductPrice, &detail.Qty)
			if err != nil {
				billRows.Close()
				return nil, err
			}

			transaction.BillDetails = append(transaction.BillDetails, detail)
		}

		billRows.Close()
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (trs *transactionRepository) GetTransactionByID(id int) (model.Transaction, error) {
	var transaction model.Transaction

	err := trs.db.QueryRow(`
		SELECT 
			t.transaction_id, 
			t.bill_date, 
			t.entry_date, 
			t.finish_date, 
			t.employee_id, 
			t.customer_id,
			COALESCE(SUM(b.product_price * b.qty), 0) AS total_bill
		FROM transaction t
		LEFT JOIN bill_details b ON t.transaction_id = b.transaction_id
		WHERE t.transaction_id = $1
		GROUP BY t.transaction_id;`, id).Scan(&transaction.Transaction_id, &transaction.BillDate, &transaction.EntryDate, &transaction.FinishDate, &transaction.EmployeeID, &transaction.CustomerID, &transaction.TotalBill)

	if err != nil {
		return model.Transaction{}, err
	}

	// Fetch Customer details using customerRepository
	customer, err := trs.customerRepository.GetCustByID(transaction.CustomerID)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction.Customer = customer

	employee, err := trs.employeeRepository.GetEmployeeByID(transaction.EmployeeID)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction.Employee = employee

	rows, err := trs.db.Query(`SELECT id, transaction_id, product_id, product_price, qty FROM bill_details WHERE transaction_id=$1;`, id)

	if err != nil {
		return model.Transaction{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var detail model.BillDetails
		err := rows.Scan(&detail.ID, &detail.BillID, &detail.ProductID, &detail.ProductPrice, &detail.Qty)
		if err != nil {
			return model.Transaction{}, err
		}
		transaction.BillDetails = append(transaction.BillDetails, detail)
	}
	return transaction, nil
}

func NewTransactionRepository(db *sql.DB, custRepo CustomerRepository, employeeRepo EmployeeRepository) TransactionRepository {
	return &transactionRepository{db: db, customerRepository: custRepo, employeeRepository: employeeRepo}
}