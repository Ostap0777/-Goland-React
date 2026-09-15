package users

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type PostgresRepository struct {
	db * sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
      log.Printf("[ERROR Create Tx Begin]: %v", err)
		return  err
	}
	defer tx.Rollback()

	UserQuery := `
		INSERT INTO users (first_name, second_name, email, password, phone, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		RETURNING id, created_at
	`

	err = r.db.QueryRowContext(
		ctx, UserQuery,
		user.FirstName,
		user.SecondName,
		user.Email,
		user.Password,
		user.Phone,
	).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}