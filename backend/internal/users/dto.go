package users

import "time"


type UserResponse struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	SecondName string   `json:"second_name"`
	Email     string    `json:"email"`
	Phone      string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	SecondName string   `json:"second_name"`
	Email     string    `json:"email"`
	Phone      string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}


func ToUserResponse(user *User) UserResponse {
	return UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Email:      user.Email,
		Phone:      user.Phone,
		CreatedAt:  user.CreatedAt,
	}
}