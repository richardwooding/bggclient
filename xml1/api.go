package xml1

import (
	"fmt"
	"github.com/richardwooding/bggclient/xml1/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Options struct {
	HttpClient *http.Client
	BaseURL    string
}

type API struct {
	httpClient *http.Client
	baseURL    string
}

func NewAPI(options Options) *API {
	return &API{
		httpClient: options.HttpClient,
		baseURL:    options.BaseURL,
	}
}

func (x *API) get(params map[string]string, elem ...string) (xmlModel model.XML1Model, err error) {
	urlStr, err := url.JoinPath(x.baseURL, elem...)
	if err != nil {
		return nil, err
	}
	url, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	values := url.Query()
	for k, v := range params {
		values.Add(k, v)
	}
	url.RawQuery = values.Encode()
	resp, err := x.httpClient.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.Decode(resp.Body)
}

type SearchOption func(m map[string]string) map[string]string

func ExactSearch() SearchOption {
	return func(m map[string]string) map[string]string {
		m["exact"] = "1"
		return m
	}
}

func (x *API) SearchBoardgames(search string, searchOptions ...SearchOption) (*model.Boardgames, error) {
	params := map[string]string{"search": search}
	for _, opt := range searchOptions {
		params = opt(params)
	}
	resp, err := x.get(params, "search")
	if err != nil {
		return nil, err
	}
	bgs, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", resp)
	}
	return bgs, nil
}

func (x *API) GetBoardgamesById(id ...string) (*model.Boardgame, error) {
	resp, err := x.get(map[string]string{}, "boardgame", strings.Join(id, ","))
	if err != nil {
		return nil, err
	}
	bg, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", resp)
	}
	return bg, nil
}
