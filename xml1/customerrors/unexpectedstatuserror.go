package customerrors

import "fmt"

type UnexpectedStatusError struct {
	Status string
}

func (e UnexpectedStatusError) Error() string {
	return fmt.Sprintf("unexpected status: %s", e.Status)
}
