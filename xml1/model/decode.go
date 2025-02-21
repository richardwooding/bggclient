package model

import (
	"encoding/xml"
	"fmt"
	"io"
)

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
			case "boardgames":
				var bgs Boardgames
				err = decoder.DecodeElement(&bgs, &token)
				result = &bgs
			case "boardgame":
				var bg Boardgame
				err = decoder.DecodeElement(&bg, &token)
				result = &bg
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
