package model

// This is for getting the individual sizes in `./views` handlers
type Sizes struct {
	Sizes []string `json:"sizes"`
}

type Product struct {
	ProductID uint32 `json:"product_id"`
	Name      string `json:"name"`
	Price     uint32 `json:"price"`
	Sizes     string `json:"sizes"`
	ImagePath string `json:"image_path"`
}
