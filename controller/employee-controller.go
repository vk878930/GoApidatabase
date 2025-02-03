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

type EmployeeController struct{
	useCase usecase.EmployeeUseCase
	rg *gin.RouterGroup
}

func (ep *EmployeeController) Route() {
	ep.rg.POST("/employees", ep.createNewEmployee)
	ep.rg.GET("/employees", ep.getAllEmployee)
	ep.rg.GET("/employees/:id", ep.getEmployeeByID)
	ep.rg.PUT("/employees",ep.updateEmployeeByID)
	ep.rg.DELETE("/employees/:id", ep.deleteEmployeeByID)
}

func (ep *EmployeeController) createNewEmployee (c *gin.Context) {
	var payload model.Employee

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	employee, err := ep.useCase.CreateNewEmployee(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (ep *EmployeeController) getAllEmployee(c *gin.Context) {
	employees, err := ep.useCase.GetAllEmployee()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve employees"})
		return
	}

	if len(employees) > 0 {
		c.JSON(http.StatusOK, employees)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "List of employees is empty"})
}

func (ep *EmployeeController) getEmployeeByID (c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employee, err := ep.useCase.GetEmployeeByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ep *EmployeeController) updateEmployeeByID(c *gin.Context){
	var payload model.Employee

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	employee, err := ep.useCase.UpdateEmployeeByID(payload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (ep *EmployeeController) deleteEmployeeByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := ep.useCase.DeleteEmployeeByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewEmployeeController (useCase usecase.EmployeeUseCase, rg *gin.RouterGroup) *EmployeeController {
	return &EmployeeController{useCase: useCase, rg:rg}
}

