package app

import "encoding/xml"

type Category struct {
	Base
	XMLName xml.Name `xml:"category"`
}

func NewCategory(id int, name string) *Category {
	return &Category{
		Base: Base{
			ID:   id,
			Name: name,
		},
	}
}
