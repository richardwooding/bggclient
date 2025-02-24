package model

import (
	"encoding/xml"
	"fmt"
	"io"
)

var BOARDGAMES_XML_NAME = "boardgames"

func Decode(reader io.Reader) (XML1Model, error) {
	decoder := xml.NewDecoder(reader)
	var result XML1Model
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch token := token.(type) {
		case xml.StartElement:
			switch token.Name.Local {
			case BOARDGAMES_XML_NAME:
				var bgs Boardgames
				err = decoder.DecodeElement(&bgs, &token)
				result = &bgs
			default:
				err = fmt.Errorf("unknown element: %v", token.Name.Local)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
