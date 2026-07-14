package model

import "encoding/xml"

type Items struct {
	XMLName    xml.Name `xml:"items" json:"-"`
	TotalItems int      `xml:"totalitems,attr" json:"totalItems"`
	TermsOfUse string   `xml:"termsofuse,attr" json:"termsOfUse"`
	PubDate    string   `xml:"pubdate,attr" json:"pubDate"`
	Items      []Item   `xml:"item" json:"items"`
}

type Item struct {
	ObjectType    string `xml:"objecttype,attr" json:"objectType"`
	ObjectID      int    `xml:"objectid,attr" json:"objectId"`
	Subtype       string `xml:"subtype,attr" json:"subtype"`
	CollID        int    `xml:"collid,attr" json:"collId"`
	Name          Name   `xml:"name" json:"name"`
	YearPublished int    `xml:"yearpublished" json:"yearPublished"`
	Image         string `xml:"image" json:"image"`
	Thumbnail     string `xml:"thumbnail" json:"thumbnail"`
	Stats         Stats  `xml:"stats" json:"stats"`
	Status        Status `xml:"status" json:"status"`
	NumPlays      int    `xml:"numplays" json:"numPlays"`
}
type Stats struct {
	MinPlayers  int    `xml:"minplayers,attr" json:"minPlayers"`
	MaxPlayers  int    `xml:"maxplayers,attr" json:"maxPlayers"`
	MinPlayTime int    `xml:"minplaytime,attr" json:"minPlayTime"`
	MaxPlayTime int    `xml:"maxplaytime,attr" json:"maxPlayTime"`
	PlayingTime int    `xml:"playingtime,attr" json:"playingTime"`
	NumOwned    int    `xml:"numowned,attr" json:"numOwned"`
	Rating      Rating `xml:"rating" json:"rating"`
}

type Rating struct {
	Value        string  `xml:"value,attr" json:"value"`
	UsersRated   int     `xml:"usersrated" json:"usersRated"`
	Average      float64 `xml:"average" json:"average"`
	BayesAverage float64 `xml:"bayesaverage" json:"bayesAverage"`
	StdDev       float64 `xml:"stddev" json:"stdDev"`
	Median       int     `xml:"median" json:"median"`
}

type Status struct {
	Own          int    `xml:"own,attr" json:"own"`
	PrevOwned    int    `xml:"prevowned,attr" json:"prevOwned"`
	ForTrade     int    `xml:"fortrade,attr" json:"forTrade"`
	Want         int    `xml:"want,attr" json:"want"`
	WantToPlay   int    `xml:"wanttoplay,attr" json:"wantToPlay"`
	WantToBuy    int    `xml:"wanttobuy,attr" json:"wantToBuy"`
	Wishlist     int    `xml:"wishlist,attr" json:"wishlist"`
	Preordered   int    `xml:"preordered,attr" json:"preordered"`
	LastModified string `xml:"lastmodified,attr" json:"lastModified"`
}
