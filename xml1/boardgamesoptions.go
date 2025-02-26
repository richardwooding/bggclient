package xml1

import (
	"errors"
	"time"
)

type BoardgameOption func(m map[string]string) (map[string]string, error)

type boardgameIncludeOption func(include bool) BoardgameOption

func boardgameParamInclude(param string) boardgameIncludeOption {
	return func(include bool) BoardgameOption {
		return func(m map[string]string) (map[string]string, error) {
			if include {
				m[param] = "1"
			} else {
				m[param] = "0"
			}
			return m, nil
		}
	}
}

var Comments = boardgameParamInclude("comments")
var Stats = boardgameParamInclude("stats")
var Historical = boardgameParamInclude("historical")

type boardgameDateOption func(date *time.Time) BoardgameOption

func boardgameParamDateOption(param string) boardgameDateOption {
	return func(date *time.Time) BoardgameOption {
		return func(m map[string]string) (map[string]string, error) {
			if date == nil {
				return nil, errors.New("date is nil")
			}
			m[param] = date.Format(time.DateOnly)
			return m, nil
		}
	}
}

var From = boardgameParamDateOption("from")
var To = boardgameParamDateOption("to")
