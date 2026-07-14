package main

import (
	"bytes"
	"encoding/json"
	"regexp"
	"testing"
)

var ansi = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func TestJSONColorizerRoundTrip(t *testing.T) {
	in := map[string]any{
		"name":   "Catan",
		"year":   1995,
		"rating": 7.1,
		"owned":  true,
		"tags":   []any{"strategy", "trading"},
		"nested": map[string]any{"empty": map[string]any{}, "list": []any{}, "null": nil},
		"quote":  `he said "hi" \ there`,
	}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	r := (&Globals{}).renderer()
	c := &jsonColorizer{w: &buf, dec: json.NewDecoder(bytes.NewReader(data)), styles: newJSONStyles(r)}
	c.dec.UseNumber()
	if err := c.value(0); err != nil {
		t.Fatalf("colorize: %v", err)
	}

	plain := ansi.ReplaceAllString(buf.String(), "")
	var out map[string]any
	if err := json.Unmarshal([]byte(plain), &out); err != nil {
		t.Fatalf("colorized output is not valid JSON: %v\n%s", err, plain)
	}
	roundTripped, _ := json.Marshal(out)
	original, _ := json.Marshal(in)
	if string(roundTripped) != string(original) {
		t.Errorf("round trip mismatch:\n got %s\nwant %s", roundTripped, original)
	}
}
