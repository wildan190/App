package app

import (
	"encoding/xml"
)

// Product adalah struct yang meng-extend Base dan mendukung XML encoding dan decoding.
type Product struct {
	Base
	XMLName  xml.Name `xml:"product" json:"-"`
	Price    int64    `json:"price" xml:"price"`
	Year     int      `json:"year" xml:"year"`
	Category Category `json:"category" xml:"category"`
}

// SetYear adalah method untuk mengatur nilai year.
func (p *Product) SetYear(year int) {
	p.Year = year
}

// GetYear adalah method untuk mendapatkan nilai year.
func (p *Product) GetYear() int {
	return p.Year
}
