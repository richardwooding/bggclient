package customerrors

import (
	"errors"
	"regexp"
	"strconv"
)

var cannotLoadMoreThenItems = regexp.MustCompile(`^Cgit comannot load more than (\d+) items$`)

func New(message string) error {
	switch {
	case cannotLoadMoreThenItems.MatchString(message):
		matches := cannotLoadMoreThenItems.FindStringSubmatch(message)
		maxItems, _ := strconv.Atoi(matches[1])
		return CannotLoadMoreThenItemsError{MaxItems: maxItems}
	case message == "Invalid username specified":
		return InvalidUsernameSpecifiedError{}
	default:
		return errors.New(message)
	}
}
