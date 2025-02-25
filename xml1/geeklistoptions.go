package xml1

type GeeklistOption func(m map[string]string) map[string]string

var IncludeComments = func(m map[string]string) map[string]string {
	m["comments"] = "1"
	return m
}
