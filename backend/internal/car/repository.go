package car

import "context"

type Repository interface {
	Create(ctx context.Context, car *Car) error
	GetByID(ctx context.Context, id int64) (*Car, error)
	List(ctx context.Context, filter Filter) ([]Car, error)
	Update(ctx context.Context, car *Car) error
	Delete(ctx context.Context, id int64) error
}
