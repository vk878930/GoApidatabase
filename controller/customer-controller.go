//layar ini berfokus untuk mendeklarasikan suatu kode untuk interaksi system lain
// layar ini mengandung HTTP method

package controller

import (
	"net/http"
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController struct{
	useCase usecase.CustUseCase
	rg *gin.RouterGroup
}

func (cu *CustomerController) Route() {
	cu.rg.POST("/customers", cu.createNewCust)
	cu.rg.GET("/customers", cu.getAllCust)
	cu.rg.GET("/customers/:id", cu.getCustByID)
	cu.rg.PUT("/customers",cu.updateCustByID)
	cu.rg.DELETE("/customers/:id", cu.deleteCustByID)
}

func (cu *CustomerController) createNewCust (c *gin.Context) {
	var payload model.Customer

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}

	customer, err := cu.useCase.CreateNewCust(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create customer data"})
		return
	}

	c.JSON(http.StatusCreated, customer)
}

func (cu *CustomerController) getAllCust(c *gin.Context) {
	customers, err := cu.useCase.GetAllCust()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve customers"})
		return
	}

	if len(customers) > 0 {
		c.JSON(http.StatusOK, customers)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List of customers is empty"})
}

func (cu *CustomerController) getCustByID (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	customer, err := cu.useCase.GetCustByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get Customer by ID"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cu *CustomerController) updateCustByID(c *gin.Context){
	var payload model.Customer

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	customer, err := cu.useCase.UpdateCustByID(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cu *CustomerController) deleteCustByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := cu.useCase.DeleteCustByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewCustController (useCase usecase.CustUseCase, rg *gin.RouterGroup) *CustomerController {
	return &CustomerController{useCase: useCase, rg:rg}
}

