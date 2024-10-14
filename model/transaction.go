package model

import (
	"fmt"
	"time"
	"encoding/json"
)

// Custom date format
const customDateFormat = "02-01-2006" // dd-MM-yyyy format

type Transaction struct {
	Transaction_id int           `json:"transaction_id"`
	BillDate      time.Time     `json:"bill_date"`
	EntryDate     time.Time     `json:"entry_date"`
	FinishDate    time.Time     `json:"finish_date"`
	EmployeeID    int           `json:"employee_id"`
	CustomerID    int           `json:"customer_id"`
	Customer      Customer      `json:"customer"`
	Employee      Employee      `json:"employee"`
	BillDetails   []BillDetails `json:"bill_details"`
	TotalBill     int           `json:"total_bill"`
}

// Custom unmarshaling to handle date in "dd-MM-yyyy" format
func (t *Transaction) UnmarshalJSON(data []byte) error {
	// Temp struct to help in the unmarshalling process
	type Alias Transaction

	aux := &struct {
		BillDate   string `json:"bill_date"`
		EntryDate  string `json:"entry_date"`
		FinishDate string `json:"finish_date"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	// Unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse the date strings into time.Time with custom format
	var err error
	if t.BillDate, err = time.Parse(customDateFormat, aux.BillDate); err != nil {
		return fmt.Errorf("invalid date format for bill_date: %v", err)
	}

	if t.EntryDate, err = time.Parse(customDateFormat, aux.EntryDate); err != nil {
		return fmt.Errorf("invalid date format for entry_date: %v", err)
	}

	if t.FinishDate, err = time.Parse(customDateFormat, aux.FinishDate); err != nil {
		return fmt.Errorf("invalid date format for finish_date: %v", err)
	}

	return nil
}


type BillDetails struct {
	ID          int    `json:"id"`          // Primary key ID of the bill details
	BillID      int    `json:"bill_id"`     // Foreign key ID of the main transaction
	ProductID   int    `json:"product_id"`  // ID of the product
	ProductName string `json:"product_name,omitempty"` // Name of the product (optional, for display purposes)
	ProductPrice int   `json:"product_price"`          // Price of the product
	Qty         int    `json:"qty"`                    // Quantity of the product in the transaction
}