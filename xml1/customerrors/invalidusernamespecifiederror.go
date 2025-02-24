package customerrors

type InvalidUsernameSpecifiedError struct {
}

func (e InvalidUsernameSpecifiedError) Error() string {
	return "invalid username specified"
}
