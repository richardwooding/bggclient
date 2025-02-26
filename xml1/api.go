package xml1

import (
	"context"
	"github.com/richardwooding/bggclient/xml1/customerrors"
	"github.com/richardwooding/bggclient/xml1/model"
	"golang.org/x/time/rate"
	"io"
	"mime"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"
)

var MAX_ALLOWED_RETRIES = 5
var MAX_ALLOWED_BOARDGAME_IDS = 20

type Options struct {
	HttpClient *http.Client
	BaseURL    string
}

type API struct {
	httpClient *http.Client
	baseURL    string
	limiter    *rate.Limiter
}

func NewAPI(options Options) *API {
	return &API{
		httpClient: options.HttpClient,
		baseURL:    options.BaseURL,
		limiter:    rate.NewLimiter(rate.Every(5), 1),
	}
}

var generalSuccessCodes = []int{http.StatusOK}
var generalRetryableCodes = []int{http.StatusTooManyRequests}
var collectrionRetryableCodes = []int{http.StatusAccepted, http.StatusTooManyRequests}
var geeklistRetryableCodes = []int{http.StatusAccepted, http.StatusTooManyRequests}

func (a *API) get(ctx context.Context, params map[string]string, successCodes []int, retryableCodes []int, elem ...string) (xmlModel model.XML1Model, err error) {
	return a.getInternal(ctx, params, successCodes, retryableCodes, MAX_ALLOWED_RETRIES, elem...)
}

func (a *API) getInternal(ctx context.Context, params map[string]string, successCodes []int, retryableCodes []int, allowedRetries int, elem ...string) (model.XML1Model, error) {
	err := a.limiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	urlStr, err := url.JoinPath(a.baseURL, elem...)
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
	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, customerrors.NotFoundError{}
	}
	if slices.Contains(retryableCodes, resp.StatusCode) {
		if allowedRetries > 0 {
			return a.getInternal(ctx, params, successCodes, retryableCodes, allowedRetries-1, elem...)
		}
		return nil, customerrors.TooManyRetriesError{Retries: MAX_ALLOWED_RETRIES}
	}
	if !slices.Contains(successCodes, resp.StatusCode) {
		contentType, _, contentTypeError := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if contentTypeError == nil {
			switch contentType {
			case "application/xml":
				_, decodedError := model.Decode(resp.Body)
				if decodedError != nil {
					return nil, decodedError
				}
			default:
				body, readError := io.ReadAll(resp.Body)
				if readError == nil {
					return nil, customerrors.New(strings.TrimSpace(string(body)))
				}
			}
		}
		return nil, customerrors.UnexpectedStatusError{Status: resp.Status}
	}
	return model.Decode(resp.Body)
}

func (x *API) SearchBoardgames(ctx context.Context, search string, searchOptions ...SearchOption) (*model.Boardgames, error) {
	params := map[string]string{"search": search}
	for _, opt := range searchOptions {
		params = opt(params)
	}
	resp, err := x.get(ctx, params, generalSuccessCodes, generalRetryableCodes, "search")
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

func (x *API) GetBoardgamesById(ctx context.Context, ids []string, options ...BoardgameOption) (*model.Boardgames, error) {
	if len(ids) > MAX_ALLOWED_BOARDGAME_IDS {
		return nil, customerrors.CannotLoadMoreThenItemsError{MaxItems: MAX_ALLOWED_BOARDGAME_IDS}
	}
	for _, id := range ids {
		if !validIdRegex.MatchString(id) {
			return nil, customerrors.InvalidIdError{ID: id}
		}
	}
	m := map[string]string{}
	var err error
	for _, option := range options {
		m, err = option(m)
		if err != nil {
			return nil, err
		}
	}
	resp, err := x.getInternal(ctx, m, generalSuccessCodes, nil, 0, "boardgame", strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	bg, ok := resp.(*model.Boardgames)
	if !ok {
		return nil, customerrors.UnexpectedResponseTypeError{Response: resp}
	}
	return bg, nil
}

func (a *API) GetBoardgameById(ctx context.Context, id string, options ...BoardgameOption) (*model.Boardgame, error) {
	bg, err := a.GetBoardgamesById(ctx, []string{id}, options...)
	if err != nil {
		return nil, err
	}
	if len(bg.Boardgames) == 0 {
		return nil, customerrors.NotFoundError{}
	}
	if bg.Boardgames[0].ObjectID != id {
		return nil, customerrors.NotFoundError{}
	}
	return &bg.Boardgames[0], nil
}

func (a *API) GetCollection(username string, collectionOptions ...CollectionOption) (*model.Items, error) {
	if username == "" {
		return nil, customerrors.InvalidUsernameSpecifiedError{}
	}
	params := map[string]string{}
	var err error
	for _, opt := range collectionOptions {
		params, err = opt(params)
		if err != nil {
			return nil, err
		}
	}
	resp, err := a.get(context.Background(), params, generalSuccessCodes, collectrionRetryableCodes, "collection", username)
	if err != nil {
		return nil, err
	}
	c, ok := resp.(*model.Items)
	if !ok {
		return nil, customerrors.UnexpectedResponseTypeError{Response: resp}
	}
	return c, nil
}

func (a *API) GetGeeklist(ctx context.Context, id string, options ...GeeklistOption) (*model.Geeklist, error) {
	if id == "" {
		return nil, customerrors.InvalidIdError{}
	}
	params := map[string]string{}
	var err error
	for _, opt := range options {
		params, err = opt(params)
		if err != nil {
			return nil, err
		}
	}
	resp, err := a.get(ctx, params, generalSuccessCodes, geeklistRetryableCodes, "geeklist", id)
	if err != nil {
		return nil, err
	}
	g, ok := resp.(*model.Geeklist)
	if !ok {
		return nil, customerrors.UnexpectedResponseTypeError{Response: resp}
	}
	return g, nil
}
