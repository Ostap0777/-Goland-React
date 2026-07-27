package car

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("car not found")
var ErrDuplicate = errors.New("car is duplicated")

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func(s *Store) List() ([]Car, error) {
	rows, err := s.db.Query(`SELECT id, make, created_at FROM cars ORDER BY name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]Car, 0)
	for rows.Next() {
		var c Car
		if err := rows.Scan(&c.ID, &c.Make, &c.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out,c)
	}
	if err := rows.Err(); err !=nil {
		return nil, err
	}
	return out, rows.Err()
}
