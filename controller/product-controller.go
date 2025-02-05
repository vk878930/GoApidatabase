package controller

import (
	"net/http"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController struct holds the usecase and router group reference
type ProductController struct {
	useCase usecase.ProductUseCase
	rg      *gin.RouterGroup
}

// Route defines the product-related HTTP routes
func (pc *ProductController) Route() {
	pc.rg.POST("/products", pc.createNewProduct)
	pc.rg.GET("/products", pc.getAllProduct)
	pc.rg.GET("/products/:id", pc.getProductByID)
	pc.rg.PUT("/products", pc.updateProductByID)
	pc.rg.DELETE("/products/:id", pc.deleteProductByID)
}

// createNewProduct handles the creation of a new product
func (pc *ProductController) createNewProduct(c *gin.Context) {
	var payload model.Product

	// Bind the JSON input to the product struct
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	product, err := pc.useCase.CreateNewProduct(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// getAllProduct handles retrieving all products from the database
func (pc *ProductController) getAllProduct(c *gin.Context) {
	products, err := pc.useCase.GetAllProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve products"})
		return
	}

	if len(products) > 0 {
		c.JSON(http.StatusOK, products)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List of products is empty"})
}

// getProductByID retrieves a specific product by its ID
func (pc *ProductController) getProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := pc.useCase.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get product by ID"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// updateProductByID updates an existing product by its ID
func (pc *ProductController) updateProductByID(c *gin.Context) {
	var payload model.Product

	// Bind the JSON input to the product struct
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	product, err := pc.useCase.UpdateProductByID(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// deleteProductByID deletes a product by its ID
func (pc *ProductController) deleteProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := pc.useCase.DeleteProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// NewProductController creates a new instance of ProductController with the given usecase and router group
func NewProductController(useCase usecase.ProductUseCase, rg *gin.RouterGroup) *ProductController {
	return &ProductController{useCase: useCase, rg: rg}
}
