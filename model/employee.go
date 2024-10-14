package model

type Employee struct {
	Employee_id int    `json:"employee_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
