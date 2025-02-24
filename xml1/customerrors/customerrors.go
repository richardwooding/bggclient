package customerrors

import "errors"

func New(message string) error {
	switch message {
	case "Invalid username specified":
		return InvalidUsernameSpecifiedError{}
	default:
		return errors.New(message)
	}
}
