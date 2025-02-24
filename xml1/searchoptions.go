package xml1

type SearchOption func(m map[string]string) map[string]string

var ExactSearch = func(m map[string]string) map[string]string {
	m["exact"] = "1"
	return m
}
