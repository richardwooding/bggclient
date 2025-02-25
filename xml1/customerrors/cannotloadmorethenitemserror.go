package customerrors

import "fmt"

type CannotLoadMoreThenItemsError struct {
	MaxItems int
}

func (e CannotLoadMoreThenItemsError) Error() string {
	return fmt.Sprintf("cannot load more then %d items", e.MaxItems)
}
