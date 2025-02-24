package model

import "encoding/xml"

type Errors struct {
	XMLName xml.Name `xml:"errors"`
	Errors  []Error  `xml:"error"`
}

type Error struct {
	Message string `xml:"message"`
}
