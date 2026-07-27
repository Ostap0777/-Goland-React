package makes

import (
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var ErrNotFound = errors.New("make not found")
var ErrDuplicate = errors.New("make already exists")

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) List() ([]Make, error) {
	rows, err := s.db.Query(`SELECT id, name, created_at FROM makes ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Make, 0)
	for rows.Next() {
		var m Make
		if err := rows.Scan(&m.ID, &m.Name, &m.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, m)
	}

	return out, rows.Err()
}

func (s *Store) Get(id int64) (Make, error) {
	var m Make
	err := s.db.QueryRow(
		`SELECT id, name, created_at FROM makes WHERE id = $1`,
		id,
	).Scan(&m.ID, &m.Name, &m.CreatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return Make{}, ErrNotFound
	}
	if err != nil {
		return Make{}, err
	}
	return m, nil
}

func (s *Store) Name(id int64) (string, bool) {
	var name string
	err := s.db.QueryRow(`SELECT name FROM makes WHERE id = $1`, id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false
	}
	return name, err == nil
}

func (s *Store) Create(in Input) (Make, error) {
	var m Make
	err := s.db.QueryRow(
		`INSERT INTO makes (name) VALUES ($1) RETURNING id, name, created_at`,
		in.Name,
	).Scan(&m.ID, &m.Name, &m.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return Make{}, ErrDuplicate
		}
		return Make{}, err
	}
	return m, nil
}

func (s *Store) Delete(id int64) error {
	res, err := s.db.Exec(`DELETE FROM makes WHERE id = $1`, id)
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
