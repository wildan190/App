// app/base.go

package app

import (
	"encoding/json"
	"encoding/xml"
)

// Base adalah struct dengan dua properti ID dan Name.
type Base struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// SetIDAndName mengimplementasikan method dari interface BaseInterface.
func (b *Base) SetIDAndName(id int, name string) {
	b.ID = id
	b.Name = name
}

// MarshalJSON customizes the JSON encoding for Base.
func (b *Base) MarshalJSON() ([]byte, error) {
	type Alias Base
	return json.Marshal(&struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		*Alias
	}{
		ID:    b.ID,
		Name:  b.Name,
		Alias: (*Alias)(b),
	})
}

// MarshalXML customizes the XML encoding for Base.
func (b *Base) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type Alias Base
	start.Name.Local = "base"
	return e.EncodeElement(&struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
		*Alias
	}{
		ID:    b.ID,
		Name:  b.Name,
		Alias: (*Alias)(b),
	}, start)
}
