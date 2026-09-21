package auth

import (
	"context"
	"errors"
	"fmt"

	"backend/internal/users"

	"golang.org/x/crypto/bcrypt"
)
var ErrEmailAlreadyExists = errors.New("user with this email already exists")
var ErrEmailNotFound = errors.New("user with this emailnot found")

type Service interface {
	Register(ctx context.Context, req RegisterRequest) (RegisterResponse, error)
	GetByEmail(ctx context.Context, email string) (RegisterResponse, error)
}

type service struct {
	userRepo users.Repository
	jwtSecret string
}

func NewService(userRepo users.Repository, jwtSecret string) Service {
	return &service{
		userRepo: userRepo,
		jwtSecret: jwtSecret,
	}
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

func (s *service) GetByEmail(ctx context.Context, email string) (RegisterResponse, error) {

	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			return RegisterResponse{}, ErrEmailNotFound
		}
		return RegisterResponse{}, fmt.Errorf("failed to get user by email: %w", err)
	}
	return RegisterResponse{
		User: users.ToUserResponse(user),
	}, nil
}