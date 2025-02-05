package repository

/*pada layar repository dia berisi query/handler
untuk database
*/

import (
	"database/sql"
	"submission-project-enigma-laundry/model"
)

type customerRepository struct {
	db *sql.DB
}

// menerapkan semua method yang berisi query kita butuhkan
type CustomerRepository interface {
	CreateNewCust(customer model.Customer) (model.Customer, error)
	GetAllCust() ([]model.Customer, error)
	GetCustByID(id int) (model.Customer, error)
	UpdateCustByID(customer model.Customer) (model.Customer, error)
	DeleteCustByID(id int) error
	IsPhoneNumberUnique(phoneNumber string) (bool, error) 

}

/*
membuat method CreateNewCust yang berhubungan dengan customerRepository struct. dia menerima model.Customer
object (merepresentasi data customer yang kita mau insert) dan mengembalikan 2 values:
model.Customer dan error
*/
func (cu *customerRepository) CreateNewCust(customer model.Customer) (model.Customer, error) {

	//mendeklarasi variable customerID tipe int. variable ini digunakan untuk menyimpan customer_id yang di return oleh SQL query setrlah insert customer baru ke database.
	var customerID int

	/*err := cu.db.QueryRow:
	cu.db adalah koneksi database (pointer ke sql.DB). kita menggunakan QueryRow method untuk eksekusi query yang mereturn sebuah row, di case ini, customer_id dari data customer yang baru kita insert.
	QueryRow digunakan untuk query yang mereturn sebuah row. karena kita insert customer baru dam mengekspektasi untuk dapat kembali customer_id, this is appropriate.*/
	err := cu.db.QueryRow(`INSERT INTO customer (name, phone, address, created_at, updated_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING customer_id;`, customer.Name, customer.Phone, customer.Address).Scan(&customerID) //scan ini digunakan untuk setor nilai RETURNING clause dari SQL query

	//mengencek error dan kembalikan mode.customer kosong dan error
	if err != nil {
		return model.Customer{}, err
	}

	//jika semua berjalan bagus nilai dari customerID yang di dapatkan dari database di assigned ke Customer_id
	customer.Customer_id = customerID

	return customer, nil
}

// method untuk mendapatkan semua buku
func (cu *customerRepository) GetAllCust() ([]model.Customer, error) {

	//menginisialisasi slice model.cust untuk menampung data dari database
	var customers []model.Customer

	//menggunakan Query daripada Queryrow karena Query untuk multiple row
	//sedangkan QueryRow hanya untuk satu row aja
	rows, err := cu.db.Query(`SELECT * FROM customer;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//mengiterasi tiap row dalam result set. rows.Next() mempersiapkan row selanjutnya untuk dibaca. dia mereturn true jika ada row lagi, atau false jika tidak ada row lagi atau  error occurred during iteration.
	for rows.Next() {
		var customer model.Customer

		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)

		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (cu *customerRepository) GetCustByID(id int) (model.Customer, error) {
	var customer model.Customer

	err := cu.db.QueryRow(`SELECT * FROM customer WHERE customer_id = $1;`, id).Scan(&customer.Customer_id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (cu *customerRepository) UpdateCustByID(customer model.Customer) (model.Customer, error) {

	//Exec function digunakan untuk SQL query karena kita tidak ekspektasi sebauh result set (cuman eksekusi update)
	_, err := cu.db.Exec(`UPDATE customer SET name = $2, phone = $3, address = $4, updated_at = CURRENT_TIMESTAMP WHERE customer_id = $1;`, customer.Customer_id, customer.Name, customer.Phone, customer.Address)

	if err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (cu *customerRepository) DeleteCustByID(id int) error {
	_, err := cu.db.Exec(`DELETE FROM customer WHERE customer_id = $1;`, id)

	if err != nil {
		return err
	}

	return nil
}

func (cu *customerRepository) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	var exists bool
	// Optimized query with EXISTS and LIMIT 1
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM customer
			WHERE phone = $1
			LIMIT 1
		);
	`
	err := cu.db.QueryRow(query, phoneNumber).Scan(&exists)
	if err != nil {
		return false, err
	}
	// If it exists, the phone number is not unique
	return !exists, nil
}


func NewCustRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}
