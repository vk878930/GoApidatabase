package repository

import (
	"database/sql"
	"submission-project-enigma-laundry/model"
	"fmt"
)

// productRepository struct implements the ProductRepository interface
type productRepository struct {
	db *sql.DB
}

// ProductRepository interface defines the methods for product-related database operations
type ProductRepository interface {
	CreateNewProduct(product model.Product) (model.Product, error)
	GetAllProduct() ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	UpdateProductByID(product model.Product) (model.Product, error)
	DeleteProductByID(id int) error
	IsProductNameUnique(name string) (bool, error) // Check if product name is unique
}

// CreateNewProduct inserts a new product into the database and returns the created product
func (pr *productRepository) CreateNewProduct(product model.Product) (model.Product, error) {
	var productID int

	// Insert product into the database
	err := pr.db.QueryRow(`
		INSERT INTO product (name, price, unit, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING product_id;`,
		product.Name, product.Price, product.Unit).Scan(&productID)

	if err != nil {
		return model.Product{}, err
	}

	product.Product_id = productID
	return product, nil
}

// GetAllProduct retrieves all products from the database
func (pr *productRepository) GetAllProduct() ([]model.Product, error) {
	var products []model.Product

	// Select all products
	rows, err := pr.db.Query(`SELECT * FROM product;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the result set
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.Product_id, &product.Name, &product.Price, &product.Unit, &product.Created_at, &product.Updated_at)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProductByID retrieves a single product by its ID from the database
func (pr *productRepository) GetProductByID(id int) (model.Product, error) {
	var product model.Product

	// Select product by ID
	err := pr.db.QueryRow(`SELECT * FROM product WHERE product_id = $1;`, id).
		Scan(&product.Product_id, &product.Name, &product.Price, &product.Unit, &product.Created_at, &product.Updated_at)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

// UpdateProductByID updates a product's details in the database based on its ID
func (pr *productRepository) UpdateProductByID(product model.Product) (model.Product, error) {
	// Update product in the database
	_, err := pr.db.Exec(`
		UPDATE product SET name = $2, price = $3, unit = $4, updated_at = CURRENT_TIMESTAMP
		WHERE product_id = $1;`,
		product.Product_id, product.Name, product.Price, product.Unit)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

// DeleteProductByID deletes a product from the database by its ID
func (pr *productRepository) DeleteProductByID(id int) error {
	_, err := pr.db.Exec(`DELETE FROM product WHERE product_id = $1;`, id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *productRepository) IsProductNameUnique(name string) (bool, error) {
	var exists bool
	// Optimized query with EXISTS and LIMIT 1
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM product
			WHERE name = $1
			LIMIT 1
		);
	`
	err := pr.db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check product name uniqueness: %w", err)
	}
	// If it exists, the name is not unique
	return !exists, nil
}

// NewProductRepository creates a new instance of productRepository with the given database connection
func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}
