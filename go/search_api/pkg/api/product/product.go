// Package product provides product related functionality.
package product

// NewProduct creates new instance of Product.
func NewProduct() *Product {
	return &Product{}
}

// Product model.
type Product struct {
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}
