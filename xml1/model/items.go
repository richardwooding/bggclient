package model

import "encoding/xml"

type Items struct {
	XMLName    xml.Name `xml:"items"`
	TotalItems int      `xml:"totalitems,attr"`
	TermsOfUse string   `xml:"termsofuse,attr"`
	PubDate    string   `xml:"pubdate,attr"`
	Items      []Item   `xml:"item"`
}

type Item struct {
	ObjectType    string `xml:"objecttype,attr"`
	ObjectID      int    `xml:"objectid,attr"`
	Subtype       string `xml:"subtype,attr"`
	CollID        int    `xml:"collid,attr"`
	Name          Name   `xml:"name"`
	YearPublished int    `xml:"yearpublished"`
	Image         string `xml:"image"`
	Thumbnail     string `xml:"thumbnail"`
	Stats         Stats  `xml:"stats"`
	Status        Status `xml:"status"`
	NumPlays      int    `xml:"numplays"`
}
type Stats struct {
	MinPlayers  int    `xml:"minplayers,attr"`
	MaxPlayers  int    `xml:"maxplayers,attr"`
	MinPlayTime int    `xml:"minplaytime,attr"`
	MaxPlayTime int    `xml:"maxplaytime,attr"`
	PlayingTime int    `xml:"playingtime,attr"`
	NumOwned    int    `xml:"numowned,attr"`
	Rating      Rating `xml:"rating"`
}

type Rating struct {
	Value        string  `xml:"value,attr"`
	UsersRated   int     `xml:"usersrated"`
	Average      float64 `xml:"average"`
	BayesAverage float64 `xml:"bayesaverage"`
	StdDev       float64 `xml:"stddev"`
	Median       int     `xml:"median"`
}

type Status struct {
	Own          int    `xml:"own,attr"`
	PrevOwned    int    `xml:"prevowned,attr"`
	ForTrade     int    `xml:"fortrade,attr"`
	Want         int    `xml:"want,attr"`
	WantToPlay   int    `xml:"wanttoplay,attr"`
	WantToBuy    int    `xml:"wanttobuy,attr"`
	Wishlist     int    `xml:"wishlist,attr"`
	Preordered   int    `xml:"preordered,attr"`
	LastModified string `xml:"lastmodified,attr"`
}
