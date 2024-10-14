package model

type Product struct {

	Product_id int    `json:"product_id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Unit       string `json:"unit"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
