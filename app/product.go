package app

import "encoding/xml"

type Product struct {
	Base
	XMLName  xml.Name  `xml:"product"`            // XMLName specifies the XML element name for the Product; ignored by JSON.
	Price    int64     `json:"price" xml:"price"` // Price of the product.
	year     int       // year is private and not accessible outside the package.
	Category *Category `json:"category" xml:"category"` // Category represents a relationship to Category.
}

// NewProduct is a constructor for creating a new Product with initial values.
func NewProduct(id int, name string, price int64, year int, category *Category) *Product {
	return &Product{
		Base: Base{
			ID:   id,
			Name: name,
		},
		Price:    price,
		year:     year,
		Category: category,
	}
}

// SetYear sets the year for the product.
func (p *Product) SetYear(year int) {
	p.year = year
}

// GetYear returns the year of the product.
func (p *Product) GetYear() int {
	return p.year
}
