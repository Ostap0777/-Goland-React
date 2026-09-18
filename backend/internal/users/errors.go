package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserDuplicate = errors.New("user already exists")
)

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}