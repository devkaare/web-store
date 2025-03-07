package model

type Product struct {
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Sizes     string `json:"sizes"`
	ImagePath string `json:"image_path"`
}
