package usecase

//usecase ini digunakan untuk setor semua business logic pada aplikasi

import (
	"fmt"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
)

// untuk menampung semua repo
type custUsecase struct {
	repo repository.CustomerRepository
}

type CustUseCase interface {
	CreateNewCust(customer model.Customer) (model.Customer, error)
	GetAllCust() ([]model.Customer, error)
	GetCustByID(id int) (model.Customer, error)
	UpdateCustByID(customer model.Customer) (model.Customer, error)
	DeleteCustByID(id int) error
}

func (cu *custUsecase) CreateNewCust(customer model.Customer) (model.Customer, error) {
	return cu.repo.CreateNewCust(customer)
}

func (cu *custUsecase) GetAllCust() ([]model.Customer, error) {
	return cu.repo.GetAllCust()
}

func (cu *custUsecase) GetCustByID(id int) (model.Customer, error) {
	return cu.repo.GetCustByID(id)
}

func (cu *custUsecase) UpdateCustByID(customer model.Customer) (model.Customer, error) {

	//karena data buku tidak pakai maka "_"
	_, err := cu.repo.GetCustByID(customer.Customer_id)

	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with ID %d not found", customer.Customer_id)
	}

	return cu.repo.UpdateCustByID(customer)
}

func (cu *custUsecase) DeleteCustByID(id int) error {
	_, err := cu.repo.GetCustByID(id)

	if err != nil {
		return fmt.Errorf("customer with ID %d not found", id)
	}

	return cu.repo.DeleteCustByID(id)
}

func NewCustUseCase(repo repository.CustomerRepository) CustUseCase {
	return &custUsecase{repo: repo}

}
