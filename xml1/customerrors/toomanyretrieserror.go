package customerrors

import "fmt"

type TooManyRetriesError struct {
	Retries int
}

func (e TooManyRetriesError) Error() string {
	return fmt.Sprintf("Too many retries: %d", e.Retries)
}
