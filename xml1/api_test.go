package xml1

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"github.com/henvic/httpretty"
	"github.com/pborman/indent"
	"github.com/richardwooding/bggclient/xml1/customerrors"
	"github.com/richardwooding/bggclient/xml1/model"
	"github.com/seborama/govcr/v15"
	"net/http"
	"os"
	"strings"
	"testing"
)

type apiKey struct{}
type resultKey struct{}
type errKey struct{}

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

	logger.SetOutput(indent.New(os.Stdout, "      [wire] "))

	httpClient := &http.Client{
		Transport: logger.RoundTripper(http.DefaultTransport),
	}

	vcr := govcr.NewVCR(
		govcr.NewCassetteLoader("fixtures/bgg.json"),
		govcr.WithClient(httpClient),
	)

	api := NewAPI(Options{
		HttpClient: vcr.HTTPClient(),
		BaseURL:    "https://boardgamegeek.com/xmlapi",
	})
	return context.WithValue(ctx, apiKey{}, api), nil
}

func iSearchFor(ctx context.Context, search string) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.SearchBoardgames(ctx, search)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
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
	return iSearchFor(ctx, "")
}

func iRequestTheBoardgameWithID(ctx context.Context, id string) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.GetBoardgameById(ctx, id)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func convertTableToStringSlice(table *godog.Table) ([]string, error) {
	var stringSlice []string
	for _, row := range table.Rows {
		for i, cell := range row.Cells {
			if i != 0 {
				return nil, errors.New("table must have only one column")
			}
			stringSlice = append(stringSlice, cell.Value)
		}
	}
	return stringSlice, nil
}

func iRequestTheBoardGamesWithIDs(ctx context.Context, table *godog.Table) (context.Context, error) {
	ids, err := convertTableToStringSlice(table)
	if err != nil {
		return ctx, err
	}
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.GetBoardgamesById(ctx, ids...)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func iRequestTheBoardgameWithAnEmptyID(ctx context.Context) (context.Context, error) {
	return iRequestTheBoardgameWithID(ctx, "")
}

func iShouldReceiveASingleBoardgame(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value(resultKey{}).(*model.Boardgame)
	if !ok {
		return ctx, errors.New("results is not boardgame")
	}
	return ctx, nil
}

func iShouldReceiveAnErrorMessage(ctx context.Context) (context.Context, error) {
	err, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	if err == nil {
		return ctx, errors.New("error is nil")
	}
	return ctx, nil
}

func theBoardgameShouldHaveTheID(ctx context.Context, expectedId string) (context.Context, error) {
	boardgame, ok := ctx.Value(resultKey{}).(*model.Boardgame)
	if !ok {
		return ctx, errors.New("result is not boardgames")
	}
	if boardgame.ObjectID != expectedId {
		return ctx, errors.New("boardgame has wrong ID")
	}
	return ctx, nil
}

func theErrorMessageShouldIndicateThatTheBoardgameWasNotFound(ctx context.Context) (context.Context, error) {
	err, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	if err == nil {
		return ctx, errors.New("error is nil")
	}
	if !errors.As(err, &customerrors.NotFoundError{}) {
		return ctx, errors.New("error is not NotFoundError")
	}
	return ctx, nil
}

func theErrorMessageShouldIndicateThatTheIDisInvalid(ctx context.Context) (context.Context, error) {
	err, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	if err == nil {
		return ctx, errors.New("error is nil")
	}
	if !strings.Contains(err.Error(), "invalid") {
		return ctx, errors.New("error message does not indicate that the ID is invalid")
	}
	return ctx, nil
}

func theListShouldContainABoardgameWithTheIDs(ctx context.Context, table *godog.Table) (context.Context, error) {
	boardgames, ok := ctx.Value(resultKey{}).(*model.Boardgames)
	if !ok {
		return ctx, errors.New("result is not boardgames")
	}
	ids, err := convertTableToStringSlice(table)
	if err != nil {
		return ctx, err
	}
	for _, id := range ids {
		found := false
		for _, bg := range boardgames.Boardgames {
			if bg.ObjectID == id {
				found = true
				break
			}
		}
		if !found {
			return ctx, errors.New("boardgame not found in results")
		}
	}
	return ctx, nil
}

func iRequestTheCollectionForUser(ctx context.Context, username string) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.GetCollection(username)
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func iShouldReceiveACollectionOfBoardgames(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value(resultKey{}).(*model.Items)
	if !ok {
		return ctx, errors.New("results is not a collection of boardgames")
	}
	return ctx, nil
}

func theCollectionShouldContainBoardgames(ctx context.Context) (context.Context, error) {
	items, ok := ctx.Value(resultKey{}).(*model.Items)
	if !ok {
		return ctx, errors.New("results is not a collection of boardgames")
	}
	if len(items.Items) == 0 {
		return ctx, errors.New("collection is empty")
	}
	return ctx, nil
}

func theErrorMessageShouldIndicateThatTheUserWasNotFound(ctx context.Context) (context.Context, error) {
	err, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	if err == nil {
		return ctx, errors.New("error is nil")
	}
	if !errors.As(err, &customerrors.InvalidUsernameSpecifiedError{}) {
		return ctx, errors.New("error is not InvalidUsernameSpecifiedError")
	}
	return ctx, nil
}

func iRequestTheCollectionWithAnEmptyUsername(ctx context.Context) (context.Context, error) {
	return iRequestTheCollectionForUser(ctx, "")
}

func theErrorMessageShouldIndicateThatTheUsernameIsInvalid(ctx context.Context) (context.Context, error) {
	err, ok := ctx.Value(errKey{}).(error)
	if !ok {
		return ctx, errors.New("error not found in context")
	}
	if err == nil {
		return ctx, errors.New("error is nil")
	}
	if !errors.As(err, &customerrors.InvalidUsernameSpecifiedError{}) {
		return ctx, errors.New("error is not InvalidUsernameSpecifiedError")
	}
	return ctx, nil
}

func iRequestTheCollectionForUserWithInt(ctx context.Context, username, filter string, value int) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.GetCollection(username, Filter(filter, value))
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func iRequestTheCollectionForUserWithFilterOn(ctx context.Context, username, filter string) (context.Context, error) {
	api, ok := ctx.Value(apiKey{}).(*API)
	if !ok {
		return ctx, errors.New("api not found in context")
	}
	results, err := api.GetCollection(username, Filter(filter, true))
	if err != nil {
		return context.WithValue(ctx, errKey{}, err), nil
	}
	return context.WithValue(ctx, resultKey{}, results), nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the API is initialized with a valid base URL and HTTP client$`, theAPIIsInitializedWithAValidBaseURLAndHTTPClient)
	ctx.Step(`^I search for "([^"]*)"$`, iSearchFor)
	ctx.Step(`^I should receive a list of boardgames$`, IShouldReceiveAListOfBoardgames)
	ctx.Step(`^the list should contain a boardgame with the name "([^"]*)"$`, theListShouldContainABoardgameWithTheName)
	ctx.Step(`^I should receive an empty list of boardgames$`, iShouldReceiveAnEmptyListOfBoardgames)
	ctx.Step(`^I search for an empty string$`, iSearchForAnEmptyString)
	ctx.Step(`^I request the boardgame with ID "([^"]*)"$`, iRequestTheBoardgameWithID)
	ctx.Step(`^I request the boardgame with an empty ID`, iRequestTheBoardgameWithAnEmptyID)
	ctx.Step(`^I request the boardgames with IDs$`, iRequestTheBoardGamesWithIDs)
	ctx.Step(`^I should receive a single boardgame$`, iShouldReceiveASingleBoardgame)
	ctx.Step(`^I should receive an error message$`, iShouldReceiveAnErrorMessage)
	ctx.Step(`^the boardgame should have the ID "([^"]*)"$`, theBoardgameShouldHaveTheID)
	ctx.Step(`^the error message should indicate that the boardgame was not found$`, theErrorMessageShouldIndicateThatTheBoardgameWasNotFound)
	ctx.Step(`^the error message should indicate that the ID is invalid$`, theErrorMessageShouldIndicateThatTheIDisInvalid)
	ctx.Step(`^the list should contain a boardgame with the IDs$`, theListShouldContainABoardgameWithTheIDs)
	ctx.Step(`^I request the collection for user "([^"]*)"$`, iRequestTheCollectionForUser)
	ctx.Step(`^I should receive a collection of boardgames$`, iShouldReceiveACollectionOfBoardgames)
	ctx.Step(`^the collection should contain boardgames$`, theCollectionShouldContainBoardgames)
	ctx.Step(`^the error message should indicate that the user was not found$`, theErrorMessageShouldIndicateThatTheUserWasNotFound)
	ctx.Step(`^I request the collection with an empty username$`, iRequestTheCollectionWithAnEmptyUsername)
	ctx.Step(`^the error message should indicate that the username is invalid$`, theErrorMessageShouldIndicateThatTheUsernameIsInvalid)
	ctx.Step(`^I request the collection for user "([^"]*)" with ([\w\s]*) (\d+)$`, iRequestTheCollectionForUserWithInt)
	ctx.Step(`^I request the collection for user "([^"]*)" with ([\w\s]*) only$`, iRequestTheCollectionForUserWithFilterOn)
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
