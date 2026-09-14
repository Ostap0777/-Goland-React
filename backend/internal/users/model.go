package users

import "time"

type User struct {
	ID         int64
	FirstName  string
	SecondName string
	Email      string
	Phone      string
	Premium    bool
	Password   string
	CreatedAt   time.Time
}