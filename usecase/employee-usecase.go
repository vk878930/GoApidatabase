package usecase

//usecase ini digunakan untuk setor semua business logic pada aplikasi

import (
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
	_, err := ep.repo.GetEmployeeByID(employee.Employee_id)

	if err != nil {
		return model.Employee{}, fmt.Errorf("employee with ID %d not found", employee.Employee_id)
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
