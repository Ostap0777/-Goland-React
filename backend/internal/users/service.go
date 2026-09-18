package users

import "context"

type Service interface {
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
func (s* service) GetByID(ctx context.Context, id int64) (*User, error) {
	return s.repo.GetByID(ctx, id)
}
func (s* service) GetByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s* service) Update(ctx context.Context, user *User) (*User, error) {

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}
	return s.repo.GetByID(ctx, user.ID)
}