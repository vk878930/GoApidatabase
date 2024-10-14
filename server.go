package main

import (
	"fmt"
	"submission-project-enigma-laundry/config"
	"submission-project-enigma-laundry/controller"
	"submission-project-enigma-laundry/repository"
	"submission-project-enigma-laundry/usecase"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	transactionUC usecase.TransactionUseCase
	productUC     usecase.ProductUseCase
	employeeUC    usecase.EmployeeUseCase
	customerUC    usecase.CustUseCase
	engine        *gin.Engine
	host          string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")

	controller.NewCustController(s.customerUC, rg).Route()
	controller.NewEmployeeController(s.employeeUC, rg).Route()
	controller.NewProductController(s.productUC, rg).Route()
	controller.NewTransactionController(s.transactionUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()

	err := s.engine.Run(s.host)

	if err != nil {
		panic(fmt.Errorf("server not running on host %s because error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	db, err := sql.Open(cfg.Driver, dsn)

	if err != nil {
		panic("connection error")
	}

	productRepo := repository.NewProductRepository(db)
	productUseCase := usecase.NewProductUseCase(productRepo)

	employeeRepo := repository.NewEmployeeRepository(db)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)

	custRepo := repository.NewCustRepository(db)
	customerUseCase := usecase.NewCustUseCase(custRepo)

	transactionRepo := repository.NewTransactionRepository(db, custRepo, employeeRepo)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		transactionUC: transactionUseCase,
		productUC:     productUseCase,
		employeeUC:    employeeUseCase,
		customerUC:    customerUseCase,
		engine:        engine,
		host:          host,
	}
}
