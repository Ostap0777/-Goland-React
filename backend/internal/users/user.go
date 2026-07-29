package users

import (
	"encoding/json"
	"errors"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"
)

type User struct {
	ID    int64 `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Blocked bool `json:"blocked"`
	
	PremiumUntil *time.Time `json:"premiumUntil"`

	PasswordHash string `json:"-"`

	CreatedAt time.Time `json:"createdAt"`
}

func (u User) IsPremium() bool {
	return  u.PremiumUntil != nil && u.PremiumUntil.After(time. Now().UTC())
}


func (u User) MarshalJSON() ([]byte, error) {
	type userJSON User
	return json.Marshal(struct {
		userJSON
		Premium bool `json:"premium"`
	}{
		userJSON: userJSON(u),
		Premium: u.IsPremium(),
	})
}
type LoginInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
}

type RegisterInput struct {
		Email string `json:"email"`
		Password string `json:"password"`
		FirstName string `json:"firstName"`
	   LastName string `json:"lastName"`
}

const (
	minPasswordRunes = 8
	maxPasswordRunes = 128
	maxFirstNameBytes = 100
	maxLastNameBytes = 100
)

//Validation for Register user
func (in *RegisterInput) Validate() error {
	email, err := canonicalEmail(in.Email)
	if err != nil {
		return nil
	}
	in.Email = email
	in.FirstName = strings.TrimSpace(in.FirstName)
	in.LastName = strings.TrimSpace(in.LastName)

	switch {
	case in.FirstName == "":
		return errors.New("firstName is required")
   case in.LastName == "":
		return errors.New("LastName is required")
   case len(in.FirstName) > maxFirstNameBytes: 
      return errors.New("FirstName is too long")
	case len(in.LastName) > maxLastNameBytes: 
	   return errors.New("Last name is too long")
	}
	return validatePassword(in.Password)
}

//Validation for Login user
func (in *LoginInput) Validate() error  {
	email, err := canonicalEmail(in.Email)
	if err != nil {
		return nil
	}
	in.Email = email

	switch {
	case in.Password == "":
		return errors.New("password is required")
	}
	return nil

}



func canonicalEmail (raw string) (string, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return "", errors.New("email is required")
	}
	addr, err := mail.ParseAddress(trimmed)

	if err != nil || addr.Name != "" {
		return "", errors.New("email is not valid")
	}
	return strings.ToLower(addr.Address), nil
}


func validatePassword(password string) error {
	switch {
	case password == "":
		return errors.New("passwors is required")
	case utf8.RuneCountInString(password) < minPasswordRunes: 
	   return errors.New("password must be at least 8 characters")
	case len(password) > maxPasswordRunes:
		return errors.New("password is too long")
	}
	return nil
}

