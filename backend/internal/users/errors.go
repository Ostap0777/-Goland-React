package users

import "errors"

var (
	ErrNotFound = errors.New("user not found")
	ErrDuplicate = errors.New("user already exists")
)

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}