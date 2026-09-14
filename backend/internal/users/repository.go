package users

import "context"

type Repository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, car *User) error
	ChangePassword(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
