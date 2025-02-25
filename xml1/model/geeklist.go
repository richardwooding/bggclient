package model

import "encoding/xml"

type Geeklist struct {
	XMLName           xml.Name       `xml:"geeklist"`
	ID                int            `xml:"id,attr"`
	TermsOfUse        string         `xml:"termsofuse,attr"`
	PostDate          string         `xml:"postdate"`
	PostDateTimestamp int            `xml:"postdate_timestamp"`
	EditDate          string         `xml:"editdate"`
	EditDateTimestamp int            `xml:"editdate_timestamp"`
	Thumbs            int            `xml:"thumbs"`
	NumItems          int            `xml:"numitems"`
	Username          string         `xml:"username"`
	Title             string         `xml:"title"`
	Description       string         `xml:"description"`
	Items             []GeeklistItem `xml:"item"`
}

type GeeklistItem struct {
	ID         int    `xml:"id,attr"`
	ObjectType string `xml:"objecttype,attr"`
	Subtype    string `xml:"subtype,attr"`
	ObjectID   int    `xml:"objectid,attr"`
	ObjectName string `xml:"objectname,attr"`
	Username   string `xml:"username,attr"`
	PostDate   string `xml:"postdate,attr"`
	EditDate   string `xml:"editdate,attr"`
	Thumbs     int    `xml:"thumbs,attr"`
	ImageID    int    `xml:"imageid,attr"`
	Body       string `xml:"body"`
}
