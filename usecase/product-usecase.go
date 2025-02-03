package usecase

import (
	"fmt"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/repository"
)

// productUsecase struct to hold the repository reference
type productUsecase struct {
	repo repository.ProductRepository
}

// ProductUseCase interface defines the business logic for product operations
type ProductUseCase interface {
	CreateNewProduct(product model.Product) (model.Product, error)
	GetAllProduct() ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	UpdateProductByID(product model.Product) (model.Product, error)
	DeleteProductByID(id int) error
}

// CreateNewProduct handles the business logic for creating a new product
func (pd *productUsecase) CreateNewProduct(product model.Product) (model.Product, error) {

	isUnique, err := pd.repo.IsProductNameUnique(product.Name)
	if err != nil {
		return model.Product{}, fmt.Errorf("error checking product name uniqueness: %v", err)
	}
	if !isUnique {
		return model.Product{}, fmt.Errorf("product name '%s' already exists", product.Name)
	}
	return pd.repo.CreateNewProduct(product)
}

// GetAllProduct retrieves a list of all products
func (pd *productUsecase) GetAllProduct() ([]model.Product, error) {
	return pd.repo.GetAllProduct()
}

// GetProductByID fetches a product by its ID from the repository
func (pd *productUsecase) GetProductByID(id int) (model.Product, error) {
	return pd.repo.GetProductByID(id)
}

// UpdateProductByID updates the details of an existing product
func (pd *productUsecase) UpdateProductByID(product model.Product) (model.Product, error) {
	// Check if the product exists before updating
	existingProduct, err := pd.repo.GetProductByID(product.Product_id)
	if err != nil {
		return model.Product{}, fmt.Errorf("product with ID %d not found", product.Product_id)
	}

	if product.Name != existingProduct.Name {
		isUnique, err := pd.repo.IsProductNameUnique(product.Name)
		if err != nil {
			return model.Product{}, fmt.Errorf("error checking product name uniqueness: %v", err)
		}
		if !isUnique {
			return model.Product{}, fmt.Errorf("product name '%s' already exists", product.Name)
		}
	}

	return pd.repo.UpdateProductByID(product)
}

// DeleteProductByID deletes a product by its ID from the repository
func (pd *productUsecase) DeleteProductByID(id int) error {
	_, err := pd.repo.GetProductByID(id)
	if err != nil {
		return fmt.Errorf("product with ID %d not found", id)
	}

	return pd.repo.DeleteProductByID(id)
}

// NewProductUseCase creates a new instance of productUsecase with the given repository
func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUsecase{repo: repo}
}
