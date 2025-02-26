package xml1

import "errors"

type GeeklistOption func(m map[string]string) (map[string]string, error)

var GeeklistComments = func(m map[string]string) (map[string]string, error) {
	m["comments"] = "1"
	return m, nil
}

func GeeklistFilter(name string) GeeklistOption {
	switch name {
	case "comments":
		return GeeklistComments
	default:
		return func(m map[string]string) (map[string]string, error) {
			return nil, errors.New("Invalid filter name specified")
		}
	}
}
