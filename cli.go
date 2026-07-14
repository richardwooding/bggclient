package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/richardwooding/bggclient/xml1"
)

type Globals struct {
	BaseURL         string           `help:"BGG XML API base URL." default:"https://boardgamegeek.com/xmlapi" env:"BGG_BASE_URL"`
	Token           string           `help:"BGG XML API token (register at https://boardgamegeek.com/using_the_xml_api)." env:"BGG_API_TOKEN"`
	RequestInterval time.Duration    `help:"Minimum interval between BGG requests." default:"5s" env:"BGG_REQUEST_INTERVAL"`
	Timeout         time.Duration    `help:"HTTP request timeout." default:"30s" env:"BGG_TIMEOUT"`
	NoColor         bool             `help:"Disable colour output (NO_COLOR is also respected)."`
	Version         kong.VersionFlag `help:"Print version and exit."`
}

var cli struct {
	Globals

	Search     SearchCmd     `cmd:"" help:"Search BoardGameGeek for board games by name."`
	Boardgame  BoardgameCmd  `cmd:"" help:"Fetch board games by BGG id."`
	Collection CollectionCmd `cmd:"" help:"Fetch a user's game collection."`
	Geeklist   GeeklistCmd   `cmd:"" help:"Fetch a geeklist by id."`
	Serve      ServeCmd      `cmd:"" help:"Run the MCP server (stdio by default)."`
}

func (g *Globals) newAPI() *xml1.API {
	return xml1.NewAPI(xml1.Options{
		HttpClient:      &http.Client{Timeout: g.Timeout},
		BaseURL:         g.BaseURL,
		APIToken:        g.Token,
		RequestInterval: g.RequestInterval,
	})
}

// renderer colours output when stdout is a colour-capable TTY. termenv's
// detection already honours the NO_COLOR standard and downgrades when
// output is not a terminal; --no-color forces plain output.
func (g *Globals) renderer() *lipgloss.Renderer {
	r := lipgloss.NewRenderer(os.Stdout)
	if g.NoColor {
		r.SetColorProfile(termenv.Ascii)
	}
	return r
}

type jsonStyles struct {
	Key    lipgloss.Style
	String lipgloss.Style
	Number lipgloss.Style
	Bool   lipgloss.Style
	Null   lipgloss.Style
	Punct  lipgloss.Style
}

func newJSONStyles(r *lipgloss.Renderer) jsonStyles {
	return jsonStyles{
		Key:    r.NewStyle().Foreground(lipgloss.Color("13")), // bright magenta
		String: r.NewStyle().Foreground(lipgloss.Color("10")), // bright green
		Number: r.NewStyle().Foreground(lipgloss.Color("11")), // bright yellow
		Bool:   r.NewStyle().Foreground(lipgloss.Color("14")), // bright cyan
		Null:   r.NewStyle().Foreground(lipgloss.Color("8")),  // grey
		Punct:  r.NewStyle().Foreground(lipgloss.Color("7")),  // light grey
	}
}

func (g *Globals) printJSON(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	r := g.renderer()
	if r.ColorProfile() == termenv.Ascii {
		var out bytes.Buffer
		if err := json.Indent(&out, data, "", "  "); err != nil {
			return err
		}
		fmt.Println(out.String())
		return nil
	}
	c := &jsonColorizer{
		w:      os.Stdout,
		dec:    json.NewDecoder(bytes.NewReader(data)),
		styles: newJSONStyles(r),
	}
	c.dec.UseNumber()
	if err := c.value(0); err != nil {
		return err
	}
	fmt.Println()
	return nil
}

type jsonColorizer struct {
	w      io.Writer
	dec    *json.Decoder
	styles jsonStyles
}

func (c *jsonColorizer) print(s string) {
	io.WriteString(c.w, s)
}

func (c *jsonColorizer) indent(level int) string {
	return strings.Repeat("  ", level)
}

// quote re-encodes s with JSON string escaping.
func quote(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (c *jsonColorizer) value(level int) error {
	tok, err := c.dec.Token()
	if err != nil {
		return err
	}
	switch t := tok.(type) {
	case json.Delim:
		switch t {
		case '{':
			return c.object(level)
		case '[':
			return c.array(level)
		}
		return fmt.Errorf("unexpected delimiter %v", t)
	case string:
		c.print(c.styles.String.Render(quote(t)))
	case json.Number:
		c.print(c.styles.Number.Render(t.String()))
	case bool:
		c.print(c.styles.Bool.Render(fmt.Sprintf("%t", t)))
	case nil:
		c.print(c.styles.Null.Render("null"))
	}
	return nil
}

func (c *jsonColorizer) object(level int) error {
	if !c.dec.More() {
		if _, err := c.dec.Token(); err != nil {
			return err
		}
		c.print(c.styles.Punct.Render("{}"))
		return nil
	}
	c.print(c.styles.Punct.Render("{") + "\n")
	first := true
	for c.dec.More() {
		if !first {
			c.print(c.styles.Punct.Render(",") + "\n")
		}
		first = false
		keyTok, err := c.dec.Token()
		if err != nil {
			return err
		}
		key, ok := keyTok.(string)
		if !ok {
			return fmt.Errorf("unexpected object key %v", keyTok)
		}
		c.print(c.indent(level+1) + c.styles.Key.Render(quote(key)) + c.styles.Punct.Render(":") + " ")
		if err := c.value(level + 1); err != nil {
			return err
		}
	}
	if _, err := c.dec.Token(); err != nil {
		return err
	}
	c.print("\n" + c.indent(level) + c.styles.Punct.Render("}"))
	return nil
}

func (c *jsonColorizer) array(level int) error {
	if !c.dec.More() {
		if _, err := c.dec.Token(); err != nil {
			return err
		}
		c.print(c.styles.Punct.Render("[]"))
		return nil
	}
	c.print(c.styles.Punct.Render("[") + "\n")
	first := true
	for c.dec.More() {
		if !first {
			c.print(c.styles.Punct.Render(",") + "\n")
		}
		first = false
		c.print(c.indent(level + 1))
		if err := c.value(level + 1); err != nil {
			return err
		}
	}
	if _, err := c.dec.Token(); err != nil {
		return err
	}
	c.print("\n" + c.indent(level) + c.styles.Punct.Render("]"))
	return nil
}
