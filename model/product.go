package model

type Product struct {
	ProductID uint32
	Name      string
	Price     uint32
	Sizes     []byte
	ImagePath string
}
