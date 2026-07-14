package model

import (
	"encoding/xml"
)

type Boardgames struct {
	XMLName    xml.Name    `xml:"boardgames" json:"-"`
	TermsOfUse string      `xml:"termsofuse,attr" json:"termsOfUse"`
	Boardgames []Boardgame `xml:"boardgame" json:"boardgames"`
}

type Boardgame struct {
	ObjectID      string               `xml:"objectid,attr" json:"objectId"`
	YearPublished int                  `xml:"yearpublished" json:"yearPublished"`
	MinPlayers    int                  `xml:"minplayers" json:"minPlayers"`
	MaxPlayers    int                  `xml:"maxplayers" json:"maxPlayers"`
	PlayingTime   int                  `xml:"playingtime" json:"playingTime"`
	MinPlayTime   int                  `xml:"minplaytime" json:"minPlayTime"`
	MaxPlayTime   int                  `xml:"maxplaytime" json:"maxPlayTime"`
	Age           int                  `xml:"age" json:"age"`
	Name          Name                 `xml:"name" json:"name"`
	Description   string               `xml:"description" json:"description"`
	Thumbnail     string               `xml:"thumbnail" json:"thumbnail"`
	Image         string               `xml:"image" json:"image"`
	Publishers    []BoardgamePublisher `xml:"boardgamepublisher" json:"publishers,omitempty"`
	Families      []BoardgameFamily    `xml:"boardgamefamily" json:"families,omitempty"`
	Categories    []BoardgameCategory  `xml:"boardgamecategory" json:"categories,omitempty"`
	Designers     []BoardgameDesigner  `xml:"boardgamedesigner" json:"designers,omitempty"`
	Artists       []BoardgameArtist    `xml:"boardgameartist" json:"artists,omitempty"`
	Expansions    []BoardgameExpansion `xml:"boardgameexpansion" json:"expansions,omitempty"`
	Polls         []Poll               `xml:"poll" json:"polls,omitempty"`
	PollSummaries []PollSummary        `xml:"poll-summary" json:"pollSummaries,omitempty"`
	Comments      []Comment            `xml:"comment" json:"comments,omitempty"`
	Statistics    *Statistics          `xml:"statistics" json:"statistics,omitempty"`
}

type Name struct {
	Primary   bool   `xml:"primary,attr" json:"primary"`
	SortIndex int    `xml:"sortindex,attr" json:"sortIndex"`
	Value     string `xml:",chardata" json:"value"`
}

type BoardgamePublisher struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Value    string `xml:",chardata" json:"value"`
}

type BoardgameFamily struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Value    string `xml:",chardata" json:"value"`
}

type BoardgameCategory struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Value    string `xml:",chardata" json:"value"`
}

type BoardgameDesigner struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Value    string `xml:",chardata" json:"value"`
}

type BoardgameArtist struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Value    string `xml:",chardata" json:"value"`
}

type BoardgameExpansion struct {
	ObjectID string `xml:"objectid,attr" json:"objectId"`
	Inbound  bool   `xml:"inbound,attr" json:"inbound"`
	Value    string `xml:",chardata" json:"value"`
}

type Poll struct {
	Name       string   `xml:"name,attr" json:"name"`
	Title      string   `xml:"title,attr" json:"title"`
	TotalVotes int      `xml:"totalvotes,attr" json:"totalVotes"`
	Results    []Result `xml:"results>result" json:"results,omitempty"`
}

type PollSummary struct {
	Name    string          `xml:"name,attr" json:"name"`
	Title   string          `xml:"title,attr" json:"title"`
	Results []SummaryResult `xml:"result" json:"results,omitempty"`
}

type Result struct {
	Value    string `xml:"value,attr" json:"value"`
	NumVotes int    `xml:"numvotes,attr" json:"numVotes"`
}

type SummaryResult struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

type Comment struct {
	Username string `xml:"username,attr" json:"username"`
	Rating   string `xml:"rating,attr" json:"rating"`
	Value    string `xml:",chardata" json:"value"`
}

type Statistics struct {
	Ratings Ratings `xml:"ratings" json:"ratings"`
}

type Ratings struct {
	UsersRated    int     `xml:"usersrated" json:"usersRated"`
	Average       float64 `xml:"average" json:"average"`
	BayesAverage  float64 `xml:"bayesaverage" json:"bayesAverage"`
	Ranks         []Rank  `xml:"ranks>rank" json:"ranks,omitempty"`
	StdDev        float64 `xml:"stddev" json:"stdDev"`
	Median        int     `xml:"median" json:"median"`
	Owned         int     `xml:"owned" json:"owned"`
	Trading       int     `xml:"trading" json:"trading"`
	Wanting       int     `xml:"wanting" json:"wanting"`
	Wishing       int     `xml:"wishing" json:"wishing"`
	NumComments   int     `xml:"numcomments" json:"numComments"`
	NumWeights    int     `xml:"numweights" json:"numWeights"`
	AverageWeight float64 `xml:"averageweight" json:"averageWeight"`
}

type Rank struct {
	Type         string        `xml:"type,attr" json:"type"`
	ID           int           `xml:"id,attr" json:"id"`
	Name         string        `xml:"name,attr" json:"name"`
	FriendlyName string        `xml:"friendlyname,attr" json:"friendlyName"`
	Value        RankedInt     `xml:"value,attr" json:"value"`
	BayesAverage RankedFloat64 `xml:"bayesaverage,attr" json:"bayesAverage"`
}
