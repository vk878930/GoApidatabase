package model

type Customer struct {
	Customer_id int    `json:"customer_id"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}
