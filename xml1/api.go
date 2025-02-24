package xml1

import (
	"context"
	"github.com/richardwooding/bggclient/xml1/customerrors"
	"github.com/richardwooding/bggclient/xml1/model"
	"net/http"
	"net/url"
	"regexp"
	"slices"
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

var generalSuccessCodes = []int{http.StatusOK}

func (x *API) get(ctx context.Context, params map[string]string, successCodes []int, elem ...string) (xmlModel model.XML1Model, err error) {
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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := x.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if !slices.Contains(successCodes, resp.StatusCode) {
		return nil, customerrors.UnexpectedStatusError{Status: resp.Status}
	}
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
	resp, err := x.get(ctx, params, generalSuccessCodes, "search")
	if err != nil {
		return nil, err
	}
	bgs, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, customerrors.UnexpectedResponseTypeError{Response: resp}
	}
	return bgs, nil
}

var validIdRegex = regexp.MustCompile(`^\d+$`)

func (x *API) GetBoardgamesById(ctx context.Context, ids ...string) (*model.Boardgames, error) {
	for _, id := range ids {
		if !validIdRegex.MatchString(id) {
			return nil, customerrors.InvalidIdError{ID: id}
		}
	}
	resp, err := x.get(ctx, map[string]string{}, generalSuccessCodes, "boardgame", strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	bg, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, customerrors.UnexpectedResponseTypeError{Response: resp}
	}
	return bg, nil
}

func (a *API) GetBoardgameById(ctx context.Context, id string) (*model.Boardgame, error) {
	bg, err := a.GetBoardgamesById(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(bg.Boardgames) == 0 {
		return nil, customerrors.NotFoundError{ID: id}
	}
	if bg.Boardgames[0].ObjectID != id {
		return nil, customerrors.NotFoundError{ID: id}
	}
	return &bg.Boardgames[0], nil
}
