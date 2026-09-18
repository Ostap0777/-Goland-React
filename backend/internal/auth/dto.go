package auth

import (
	"backend/internal/users"
)


type RegisterRequest struct {
   FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
}

type RegisterResponse struct {
   User  users.UserResponse `json:"user"`
	Token string             `json:"token,omitempty"`
}