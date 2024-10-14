package repository

/*pada layar repository dia berisi query/handler
untuk database
*/

import (
	"database/sql"
	"submission-project-enigma-laundry/model"
)

type employeeRepository struct {
	db *sql.DB
}

type EmployeeRepository interface {
	CreateNewEmployee(employee model.Employee) (model.Employee, error)
	GetAllEmployee() ([]model.Employee, error)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployeeByID(employee model.Employee) (model.Employee, error)
	DeleteEmployeeByID(id int) error
}


func (ep *employeeRepository) CreateNewEmployee(employee model.Employee) (model.Employee, error) {

	var employeeID int

	err := ep.db.QueryRow(`INSERT INTO employee (name, phone, address, created_at, updated_at)
	VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	RETURNING employee_id;`, employee.Name, employee.Phone, employee.Address).Scan(&employeeID) 

	if err != nil {
		return model.Employee{}, err
	}

	employee.Employee_id = employeeID

	return employee, nil
}

func (ep *employeeRepository) GetAllEmployee() ([]model.Employee, error) {

	var employees []model.Employee

	
	rows, err := ep.db.Query(`SELECT * FROM employee;`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee model.Employee

		err := rows.Scan(&employee.Employee_id, &employee.Name, &employee.Phone, &employee.Address, &employee.Created_at, &employee.Updated_at)

		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}
	return employees, nil
}

func (ep *employeeRepository) GetEmployeeByID(id int) (model.Employee, error) {
	var employee model.Employee

	err := ep.db.QueryRow(`SELECT * FROM employee WHERE employee_id = $1;`, id).Scan(&employee.Employee_id, &employee.Name, &employee.Phone, &employee.Address, &employee.Created_at, &employee.Updated_at)

	if err != nil {
		return model.Employee{}, err
	}

	return employee, nil
}

func (ep *employeeRepository) UpdateEmployeeByID(employee model.Employee) (model.Employee, error) {

	_, err := ep.db.Exec(`UPDATE employee SET name = $2, phone = $3, address = $4, updated_at = CURRENT_TIMESTAMP WHERE employee_id = $1;`, employee.Employee_id, employee.Name, employee.Phone, employee.Address)

	if err != nil {
		return model.Employee{}, err
	}

	return employee, nil
}

func (ep *employeeRepository) DeleteEmployeeByID(id int) error {
	_, err := ep.db.Exec(`DELETE FROM employee WHERE employee_id = $1;`, id)

	if err != nil {
		return err
	}

	return nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
