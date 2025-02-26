package model

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

var NOT_RANKED = "Not Ranked"

type RankedInt struct {
	Ranked bool
	Value  *int
}

func (a *RankedInt) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	default:
		v, err := strconv.Atoi(attr.Value)
		if err != nil {
			return err
		}
		a.Value = &v
		a.Ranked = true
	case NOT_RANKED:
		a.Ranked = false
		a.Value = nil
	}
	return nil
}

func (a RankedInt) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if a.Ranked {
		return xml.Attr{
			Name:  name,
			Value: strconv.Itoa(*a.Value),
		}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: NOT_RANKED,
		}, nil
	}
}

type RankedFloat64 struct {
	Ranked bool
	Value  *float64
}

func (a *RankedFloat64) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	default:
		var err error
		*a.Value, err = strconv.ParseFloat(attr.Value, 64)
		if err != nil {
			return err
		}
		a.Ranked = true
	case NOT_RANKED:
		a.Ranked = false
		a.Value = nil
	}
	return nil
}

func (a RankedFloat64) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if a.Ranked {
		return xml.Attr{
			Name:  name,
			Value: fmt.Sprintf("%f", *a.Value),
		}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: NOT_RANKED,
		}, nil
	}
}
