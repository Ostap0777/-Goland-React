package users

import (
	"strings"
)

func (in *User) Validate() error {
	in.FirstName = strings.TrimSpace(in.FirstName)
	in.SecondName = strings.TrimSpace(in.SecondName)
	in.Email = strings.TrimSpace(in.Email)

	switch {
	case in.FirstName == "":
		return &InputError{Message: "FirstName is required"}
	case in.SecondName == "":
		return &InputError{Message: "SecondName is required"}
	case in.Email == "":
		return &InputError{Message: "Email is required"}
   case len(in.Password) < 4:
	return &InputError{Message: "at least 4  is required"}
	   case len(in.Password) > 20:
	return &InputError{Message: "maximum 20 images allowed"}
		}
	return nil
}
