package usecase

//usecase ini digunakan untuk setor semua business logic pada aplikasi

import (
	"errors"
	"fmt"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
)

// untuk menampung semua repo
type employeeUsecase struct {
	repo repository.EmployeeRepository
}

type EmployeeUseCase interface {
	CreateNewEmployee(employee model.Employee) (model.Employee, error)
	GetAllEmployee() ([]model.Employee, error)
	GetEmployeeByID(id int) (model.Employee, error)
	UpdateEmployeeByID(employee model.Employee) (model.Employee, error)
	DeleteEmployeeByID(id int) error
}

func (ep *employeeUsecase) CreateNewEmployee(employee model.Employee) (model.Employee, error) {

	isUnique, err := ep.repo.IsPhoneNumberUnique(employee.Phone)
	if err != nil {
		return model.Employee{}, fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if !isUnique {
		return model.Employee{}, errors.New("phone number already exists")
	}

	return ep.repo.CreateNewEmployee(employee)
}

func (ep *employeeUsecase) GetAllEmployee() ([]model.Employee, error) {
	return ep.repo.GetAllEmployee()
}

func (ep *employeeUsecase) GetEmployeeByID(id int) (model.Employee, error) {
	
	return ep.repo.GetEmployeeByID(id)
}

func (ep *employeeUsecase) UpdateEmployeeByID(employee model.Employee) (model.Employee, error) {

	//karena data buku tidak pakai maka "_"
	existingEmployee, err := ep.repo.GetEmployeeByID(employee.Employee_id)

	if err != nil {
		return model.Employee{}, fmt.Errorf("employee with ID %d not found", employee.Employee_id)
	}

	// Check if the phone number is being updated
	if employee.Phone != existingEmployee.Phone {
		isUnique, err := ep.repo.IsPhoneNumberUnique(employee.Phone)
		if err != nil {
			return model.Employee{}, fmt.Errorf("error checking phone number uniqueness: %v", err)
		}
		if !isUnique {
			return model.Employee{}, errors.New("phone number already exists")
		}
	}
	return ep.repo.UpdateEmployeeByID(employee)
}

func (ep *employeeUsecase) DeleteEmployeeByID(id int) error {
	_, err := ep.repo.GetEmployeeByID(id)

	if err != nil {
		return fmt.Errorf("employee with ID %d not found", id)
	}

	return ep.repo.DeleteEmployeeByID(id)
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeeUseCase {
	return &employeeUsecase{repo: repo}

}
