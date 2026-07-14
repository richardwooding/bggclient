package model

import "encoding/xml"

type Geeklist struct {
	XMLName           xml.Name          `xml:"geeklist" json:"-"`
	ID                int               `xml:"id,attr" json:"id"`
	TermsOfUse        string            `xml:"termsofuse,attr" json:"termsOfUse"`
	PostDate          string            `xml:"postdate" json:"postDate"`
	PostDateTimestamp int               `xml:"postdate_timestamp" json:"postDateTimestamp"`
	EditDate          string            `xml:"editdate" json:"editDate"`
	EditDateTimestamp int               `xml:"editdate_timestamp" json:"editDateTimestamp"`
	Thumbs            int               `xml:"thumbs" json:"thumbs"`
	NumItems          int               `xml:"numitems" json:"numItems"`
	Username          string            `xml:"username" json:"username"`
	Title             string            `xml:"title" json:"title"`
	Description       string            `xml:"description" json:"description"`
	Comments          []GeeklistComment `xml:"comment" json:"comments,omitempty"`
	Items             []GeeklistItem    `xml:"item" json:"items"`
}

type GeeklistItem struct {
	ID         int    `xml:"id,attr" json:"id"`
	ObjectType string `xml:"objecttype,attr" json:"objectType"`
	Subtype    string `xml:"subtype,attr" json:"subtype"`
	ObjectID   int    `xml:"objectid,attr" json:"objectId"`
	ObjectName string `xml:"objectname,attr" json:"objectName"`
	Username   string `xml:"username,attr" json:"username"`
	PostDate   string `xml:"postdate,attr" json:"postDate"`
	EditDate   string `xml:"editdate,attr" json:"editDate"`
	Thumbs     int    `xml:"thumbs,attr" json:"thumbs"`
	ImageID    int    `xml:"imageid,attr" json:"imageId"`
	Body       string `xml:"body" json:"body"`
}

type GeeklistComment struct {
	Username string `xml:"username,attr" json:"username"`
	Date     string `xml:"date,attr" json:"date"`
	PostDate string `xml:"postdate,attr" json:"postDate"`
	EditDate string `xml:"editdate,attr" json:"editDate"`
	Thumbs   int    `xml:"thumbs,attr" json:"thumbs"`
	Body     string `xml:",chardata" json:"body"`
}
