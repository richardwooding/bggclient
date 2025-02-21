package xml1

import (
	"context"
	"fmt"
	"github.com/richardwooding/bggclient/xml1/model"
	"net/http"
	"net/url"
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

func (x *API) get(cfx context.Context, params map[string]string, elem ...string) (xmlModel model.XML1Model, err error) {
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
	req, err := http.NewRequestWithContext(cfx, http.MethodGet, url.String(), nil)
	resp, err := x.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return model.Decode(resp.Body)
}

type SearchOption func(m map[string]string) map[string]string

var ExactSearch = func(m map[string]string) map[string]string {
	m["exact"] = "1"
	return m
}

func (x *API) SearchBoardgames(ctx context.Context, search string, searchOptions ...SearchOption) (*model.Boardgames, error) {
	params := map[string]string{"search": search}
	for _, opt := range searchOptions {
		params = opt(params)
	}
	resp, err := x.get(ctx, params, "search")
	if err != nil {
		return nil, err
	}
	bgs, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", resp)
	}
	return bgs, nil
}

func (x *API) GetBoardgamesById(ctx context.Context, ids ...string) (*model.Boardgames, error) {
	resp, err := x.get(ctx, map[string]string{}, "boardgame", strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	bg, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, fmt.Errorf("unexpected type: %T", resp)
	}
	return bg, nil
}
