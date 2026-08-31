package car

import "context"

const (
	defaultListLimit = 10
	maxListLimit     = 100
)

type Service interface {
	CreateCar(ctx context.Context, req CreateCarRequest) (*Car, error)
	GetCar(ctx context.Context, id int64) (*Car, error)
	ListCars(ctx context.Context, filter Filter) ([]Car, error)
	UpdateCar(ctx context.Context, id int64, req CreateCarRequest) (*Car, error)
	DeleteCar(ctx context.Context, id int64) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCar(ctx context.Context, req CreateCarRequest) (*Car, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	car := requestToCar(0, req)
	if err := s.repo.Create(ctx, car); err != nil {
		return nil, err
	}
	return car, nil
}

func (s *service) GetCar(ctx context.Context, id int64) (*Car, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) ListCars(ctx context.Context, filter Filter) ([]Car, error) {
	if filter.Limit <= 0 {
		filter.Limit = defaultListLimit
	}
	if filter.Limit > maxListLimit {
		filter.Limit = maxListLimit
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}
	return s.repo.List(ctx, filter)
}

func (s *service) UpdateCar(ctx context.Context, id int64, req CreateCarRequest) (*Car, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	car := requestToCar(id, req)
	if err := s.repo.Update(ctx, car); err != nil {
		return nil, err
	}
	return car, nil
}

func (s *service) DeleteCar(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func requestToCar(id int64, req CreateCarRequest) *Car {
	return &Car{
		ID:          id,
		Make:        req.Make,
		Model:       req.Model,
		Year:        req.Year,
		Price:       req.Price,
		Mileage:     req.Mileage,
		Description: req.Description,
		SellerName:  req.SellerName,
	}
}
