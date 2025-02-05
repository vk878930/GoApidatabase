package controller

import (
	"submission-project-enigma-laundry/model"
	"submission-project-enigma-laundry/usecase"
	"submission-project-enigma-laundry/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	useCase usecase.UserUseCase
	rg      *gin.RouterGroup
	authMiddleware 	middleware.AuthMiddleware
}

func (uc *UserController) createUser(c *gin.Context) {
	var payload model.UserCredential
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user, err := uc.useCase.RegisterNewUser(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) getAllUser(c *gin.Context) {
	users, err := uc.useCase.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to retrieve data users"})
		return
	}

	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "List user empty"})
}

func (uc *UserController) getUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.useCase.FindUserById(uint32(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "Failed to get user by ID"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Route() {
	uc.rg.POST("/users", uc.authMiddleware.RequireToken("admin"), uc.createUser)
	uc.rg.GET("/users", uc.authMiddleware.RequireToken("admin"), uc.getAllUser)
	uc.rg.GET("/users/:id", uc.authMiddleware.RequireToken("admin"), uc.getUserById)
}

func NewUserController(useCase usecase.UserUseCase, rg *gin.RouterGroup, am middleware.AuthMiddleware) *UserController {
	return &UserController{useCase: useCase, rg: rg, authMiddleware: am }
}
