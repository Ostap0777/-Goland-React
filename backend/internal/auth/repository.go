package auth

import (
	"context"

	"backend/internal/users"
)

type Repository interface {
	CreateUser(ctx context.Context, user *users.User) error
	GetByEmail(ctx context.Context, email string) (*users.User, error)
}