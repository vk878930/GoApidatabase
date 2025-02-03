package usecase

//usecase ini digunakan untuk setor semua business logic pada aplikasi

import (
	"errors"
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
	// Check if the phone number is unique
	isUnique, err := cu.repo.IsPhoneNumberUnique(customer.Phone)
	if err != nil {
		return model.Customer{}, fmt.Errorf("error checking phone number uniqueness: %v", err)
	}
	if !isUnique {
		return model.Customer{}, errors.New("phone number already exists")
	}

	return cu.repo.CreateNewCust(customer)
}

func (cu *custUsecase) GetAllCust() ([]model.Customer, error) {
	return cu.repo.GetAllCust()
}

func (cu *custUsecase) GetCustByID(id int) (model.Customer, error) {
	return cu.repo.GetCustByID(id)
}

func (cu *custUsecase) UpdateCustByID(customer model.Customer) (model.Customer, error) {

	existingCustomer, err := cu.repo.GetCustByID(customer.Customer_id)

	if err != nil {
		return model.Customer{}, fmt.Errorf("customer with ID %d not found", customer.Customer_id)
	}

	// Check if the phone number is being updated
	if customer.Phone != existingCustomer.Phone {
		isUnique, err := cu.repo.IsPhoneNumberUnique(customer.Phone)
		if err != nil {
			return model.Customer{}, fmt.Errorf("error checking phone number uniqueness: %v", err)
		}
		if !isUnique {
			return model.Customer{}, errors.New("phone number already exists")
		}
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
