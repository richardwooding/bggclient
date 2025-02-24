package customerrors

import "fmt"

type InvalidIdError struct {
	ID string
}

func (e InvalidIdError) Error() string {
	return fmt.Sprint("invalid ID: ", e.ID)
}
