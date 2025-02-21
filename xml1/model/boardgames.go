package model

import (
	"encoding/xml"
)

type Boardgames struct {
	XMLName    xml.Name    `xml:"boardgames"`
	TermsOfUse string      `xml:"termsofuse,attr"`
	Boardgames []Boardgame `xml:"boardgame"`
}

type Boardgame struct {
	ObjectID      string               `xml:"objectid,attr"`
	YearPublished int                  `xml:"yearpublished"`
	MinPlayers    int                  `xml:"minplayers"`
	MaxPlayers    int                  `xml:"maxplayers"`
	PlayingTime   int                  `xml:"playingtime"`
	MinPlayTime   int                  `xml:"minplaytime"`
	MaxPlayTime   int                  `xml:"maxplaytime"`
	Age           int                  `xml:"age"`
	Name          Name                 `xml:"name"`
	Description   string               `xml:"description"`
	Thumbnail     string               `xml:"thumbnail"`
	Image         string               `xml:"image"`
	Publishers    []BoardgamePublisher `xml:"boardgamepublisher"`
	Families      []BoardgameFamily    `xml:"boardgamefamily"`
	Categories    []BoardgameCategory  `xml:"boardgamecategory"`
	Designers     []BoardgameDesigner  `xml:"boardgamedesigner"`
	Artists       []BoardgameArtist    `xml:"boardgameartist"`
	Expansions    []BoardgameExpansion `xml:"boardgameexpansion"`
	Polls         []Poll               `xml:"poll"`
	PollSummaries []PollSummary        `xml:"poll-summary"`
}

type Name struct {
	Primary   bool   `xml:"primary,attr"`
	SortIndex int    `xml:"sortindex,attr"`
	Value     string `xml:",chardata"`
}

type BoardgamePublisher struct {
	ObjectID string `xml:"objectid,attr"`
	Value    string `xml:",chardata"`
}

type BoardgameFamily struct {
	ObjectID string `xml:"objectid,attr"`
	Value    string `xml:",chardata"`
}

type BoardgameCategory struct {
	ObjectID string `xml:"objectid,attr"`
	Value    string `xml:",chardata"`
}

type BoardgameDesigner struct {
	ObjectID string `xml:"objectid,attr"`
	Value    string `xml:",chardata"`
}

type BoardgameArtist struct {
	ObjectID string `xml:"objectid,attr"`
	Value    string `xml:",chardata"`
}

type BoardgameExpansion struct {
	ObjectID string `xml:"objectid,attr"`
	Inbound  bool   `xml:"inbound,attr"`
	Value    string `xml:",chardata"`
}

type Poll struct {
	Name       string   `xml:"name,attr"`
	Title      string   `xml:"title,attr"`
	TotalVotes int      `xml:"totalvotes,attr"`
	Results    []Result `xml:"results>result"`
}

type PollSummary struct {
	Name    string          `xml:"name,attr"`
	Title   string          `xml:"title,attr"`
	Results []SummaryResult `xml:"result"`
}

type Result struct {
	Value    string `xml:"value,attr"`
	NumVotes int    `xml:"numvotes,attr"`
}

type SummaryResult struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
