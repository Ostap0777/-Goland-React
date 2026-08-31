package car

import "errors"

var (
	ErrNotFound    = errors.New("car not found")
	ErrUnknownMake = errors.New("make not found")
)

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}
