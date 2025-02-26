package xml1

import (
	"errors"
	"regexp"
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

func boardGameFilterBool(name string, b bool) BoardgameOption {
	switch name {
	case "comments":
		return Comments(b)
	case "stats":
		return Stats(b)
	case "history", "historical", "historical data":
		return Historical(b)
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, errors.New("Invalid filter name specified")
		}
	}
}

func boardGameFilterTime(name string, t *time.Time) BoardgameOption {
	switch name {
	case "from":
		return From(t)
	case "to":
		return To(t)
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, errors.New("Invalid filter name specified")
		}
	}
}

var datePattern = regexp.MustCompile(`\d+\-\d+\-\d+`)

func BoardgameFilter(name string, value any) BoardgameOption {
	switch v := value.(type) {
	case bool:
		return boardGameFilterBool(name, v)
	case *time.Time:
		return boardGameFilterTime(name, v)
	case string:
		if datePattern.MatchString(v) {
			t, err := time.Parse(time.DateOnly, v)
			if err != nil {
				return func(m map[string]string) (map[string]string, error) {
					return nil, err
				}
			}
			return boardGameFilterTime(name, &t)
		}
	}
	return func(m map[string]string) (map[string]string, error) {
		return nil, errors.New("Invalid filter value specified")
	}
}
