package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

const carSelect = `
SELECT c.id, m.name, c.model, c.year, c.price, c.mileage, c.description, c.seller_name, c.created_at
FROM cars c
JOIN makes m ON m.id = c.make_id`

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, car *Car) error {
	query := `
INSERT INTO cars (make_id, model, year, price, mileage, description, seller_name)
SELECT m.id, $2, $3, $4, $5, $6, $7
FROM makes m
WHERE LOWER(m.name) = LOWER($1)
RETURNING id, created_at`

	err := r.db.QueryRowContext(
		ctx, query,
		car.Make, car.Model, car.Year, car.Price, car.Mileage, car.Description, car.SellerName,
	).Scan(&car.ID, &car.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrUnknownMake
	}
	return err
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (*Car, error) {
	var c Car
	err := r.db.QueryRowContext(ctx, carSelect+` WHERE c.id = $1`, id).Scan(
		&c.ID, &c.Make, &c.Model, &c.Year, &c.Price, &c.Mileage, &c.Description, &c.SellerName, &c.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *PostgresRepository) List(ctx context.Context, filter Filter) ([]Car, error) {
	var (
		b    strings.Builder
		args []any
	)
	b.WriteString(carSelect)
	b.WriteString(" WHERE 1=1")

	add := func(sqlExpr string, value any) {
		args = append(args, value)
		fmt.Fprintf(&b, " AND "+sqlExpr, len(args))
	}

	if filter.Make != "" {
		add("LOWER(m.name) = LOWER($%d)", filter.Make)
	}
	if filter.Model != "" {
		add("LOWER(c.model) = LOWER($%d)", filter.Model)
	}
	if filter.MinYear > 0 {
		add("c.year >= $%d", filter.MinYear)
	}
	if filter.MaxYear > 0 {
		add("c.year <= $%d", filter.MaxYear)
	}
	if filter.MinPrice > 0 {
		add("c.price >= $%d", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		add("c.price <= $%d", filter.MaxPrice)
	}

	args = append(args, filter.Limit, filter.Offset)
	fmt.Fprintf(&b, " ORDER BY c.created_at DESC LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.QueryContext(ctx, b.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Car, 0)
	for rows.Next() {
		var c Car
		if err := rows.Scan(
			&c.ID, &c.Make, &c.Model, &c.Year, &c.Price, &c.Mileage, &c.Description, &c.SellerName, &c.CreatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

func (r *PostgresRepository) Update(ctx context.Context, car *Car) error {
	query := `
UPDATE cars c
SET make_id = m.id,
    model = $3,
    year = $4,
    price = $5,
    mileage = $6,
    description = $7,
    seller_name = $8
FROM makes m
WHERE c.id = $1 AND LOWER(m.name) = LOWER($2)
RETURNING c.id, m.name, c.model, c.year, c.price, c.mileage, c.description, c.seller_name, c.created_at`

	err := r.db.QueryRowContext(
		ctx, query,
		car.ID, car.Make, car.Model, car.Year, car.Price, car.Mileage, car.Description, car.SellerName,
	).Scan(
		&car.ID, &car.Make, &car.Model, &car.Year, &car.Price, &car.Mileage, &car.Description, &car.SellerName, &car.CreatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		exists, lookupErr := r.exists(ctx, car.ID)
		if lookupErr != nil {
			return lookupErr
		}
		if !exists {
			return ErrNotFound
		}
		return ErrUnknownMake
	}
	return err
}

func (r *PostgresRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, `DELETE FROM cars WHERE id = $1`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *PostgresRepository) exists(ctx context.Context, id int64) (bool, error) {
	var found bool
	err := r.db.QueryRowContext(ctx, `SELECT EXISTS(SELECT 1 FROM cars WHERE id = $1)`, id).Scan(&found)
	return found, err
}
