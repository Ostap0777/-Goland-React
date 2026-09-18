package auth

import (
	"context"
	"errors"
	"fmt"

	"backend/internal/users"

	"golang.org/x/crypto/bcrypt"
)
var ErrEmailAlreadyExists = errors.New("user with this email already exists")

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error)
}

type service struct {
	userRepo users.Repository // або auth.Repository
}

func NewService(userRepo users.Repository) Service {
	return &service{userRepo: userRepo}
}

func (s *service) Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &users.User{
		FirstName:  req.FirstName,
		SecondName: req.SecondName,
		Email:      req.Email,
		Password:   string(hashedPassword),
		Phone:      req.Phone,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return RegisterResponse{}, err
	}

	return RegisterResponse{
		User: users.ToUserResponse(user),
	}, nil
}