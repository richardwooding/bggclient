package customerrors

import "fmt"

type UnexpectedResponseTypeError struct {
	Response any
}

func (e UnexpectedResponseTypeError) Error() string {
	return fmt.Sprintf("unexpected response type: %T", e.Response)
}
