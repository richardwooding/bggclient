package xml1

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/henvic/httpretty"
	"github.com/richardwooding/bggclient/xml1/model"
	"net/http"
	"strings"
	"testing"
)

type apiKey struct{}
type resultKey struct{}

func theAPIIsInitializedWithAValidBaseURLAndHTTPClient(ctx context.Context) (context.Context, error) {

	logger := &httpretty.Logger{
		Time:           true,
		TLS:            true,
		RequestHeader:  true,
		RequestBody:    true,
		ResponseHeader: true,
		ResponseBody:   true,
		Colors:         true, // erase line if you don't like colors
		Formatters:     []httpretty.Formatter{&httpretty.JSONFormatter{}},
	}

	api := NewAPI(Options{
		HttpClient: &http.Client{
			Transport: logger.RoundTripper(http.DefaultTransport),
		},
		BaseURL: "https://boardgamegeek.com/xmlapi",
	})
	return context.WithValue(ctx, apiKey{}, api), nil
}

func ISearchFor(ctx context.Context, search string) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.SearchBoardgames(search)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func IShouldReceiveAListOfBoardgames(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value(resultKey{}).(*model.Boardgames)
	if !ok {
		return ctx, errors.New("results is not a list of boardgames")
	}
	return ctx, nil
}

func theListShouldContainABoardgameWithTheName(ctx context.Context, boardgame string) (context.Context, error) {
	result, ok := ctx.Value(resultKey{}).(*model.Boardgames)
	if !ok {
		return ctx, errors.New("results is not a list of boardgames")
	}
	for _, result := range result.Boardgames {
		if strings.Contains(strings.ToLower(result.Name.Value), strings.ToLower(boardgame)) {
			return ctx, nil
		}
	}
	return ctx, errors.New("boardgame not found in results")
}

func iShouldReceiveAnEmptyListOfBoardgames(ctx context.Context) (context.Context, error) {
	results, ok := ctx.Value(resultKey{}).(*model.Boardgames)
	if !ok {
		return ctx, errors.New("results is not a list of boardgames")
	}
	if len(results.Boardgames) != 0 {
		return ctx, errors.New("results is not empty")
	}
	return ctx, nil
}

func iSearchForAnEmptyString(ctx context.Context) (context.Context, error) {
	return ISearchFor(ctx, "")
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the API is initialized with a valid base URL and HTTP client$`, theAPIIsInitializedWithAValidBaseURLAndHTTPClient)
	ctx.Step(`^I search for "([^"]*)"$`, ISearchFor)
	ctx.Step(`^I should receive a list of boardgames$`, IShouldReceiveAListOfBoardgames)
	ctx.Step(`^the list should contain a boardgame with the name "([^"]*)"$`, theListShouldContainABoardgameWithTheName)
	ctx.Step(`^I should receive an empty list of boardgames$`, iShouldReceiveAnEmptyListOfBoardgames)
	ctx.Step(`^I search for an empty string$`, iSearchForAnEmptyString)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
