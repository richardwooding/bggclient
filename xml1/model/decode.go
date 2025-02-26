package model

import (
	"encoding/xml"
	"fmt"
	"github.com/richardwooding/bggclient/xml1/customerrors"
	"io"
)

var BOARDGAMES_XML_NAME = "boardgames"
var GEEKLIST_XML_NAME = "geeklist"
var ITEMS_XML_NAME = "items"
var ERRORS_XML_NAME = "errors"

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
			case ERRORS_XML_NAME:
				var errors Errors
				err = decoder.DecodeElement(&errors, &token)
				result = &errors
			case BOARDGAMES_XML_NAME:
				var bgs Boardgames
				err = decoder.DecodeElement(&bgs, &token)
				result = &bgs
			case ITEMS_XML_NAME:
				var items Items
				err = decoder.DecodeElement(&items, &token)
				result = &items
			case GEEKLIST_XML_NAME:
				var geeklist Geeklist
				err = decoder.DecodeElement(&geeklist, &token)
				result = &geeklist
			default:
				err = fmt.Errorf("unknown element: %v", token.Name.Local)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if errors, ok := result.(*Errors); ok {
		return nil, customerrors.New(errors.Errors[0].Message)
	}
	return result, nil
}
