package app

import (
	"encoding/xml"
)

// Category adalah struct yang meng-extend Base dan mendukung XML encoding dan decoding.
type Category struct {
	Base
	XMLName xml.Name `xml:"category" json:"-"`
}
