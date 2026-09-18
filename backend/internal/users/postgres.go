package users

import (
	"context"
	"database/sql"
	"errors"
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

	return tx.Commit()
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
	SELECT first_name, second_name, email, password, phone, created_at
	FROM users
	WHERE id = $1
	`
	user := &User{}
	err := r.db.QueryRowContext(ctx,query, id).Scan(
      &user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}
	if err != nil {
	return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return user, nil
}



func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
	SELECT first_name, second_name, email, password, phone, created_at
	FROM users
	WHERE email = $1
	`
	user := &User{}
	err := r.db.QueryRowContext(ctx,query, email).Scan(
      &user.ID,
		&user.FirstName,
		&user.SecondName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrUserNotFound
	}
	if err != nil {
	return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}


func (r *PostgresRepository) Update(ctx context.Context, user *User) error {
	query := `
	UPDATE users
	SET first_name = $1, second_name = $2, phone = $3
	WHERE id = $4
	`

	result, err := r.db.ExecContext(ctx, query, user.FirstName, user.SecondName, user.Phone, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrUserNotFound
	}
	return nil

}