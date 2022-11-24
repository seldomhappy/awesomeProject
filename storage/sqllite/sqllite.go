package sqllite

import (
	. "awesomwProject/storage"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var _ IStorage = new(Storage)

type Storage struct {
	db *sql.DB
}

// New creates new SQLite storage.
func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Store(ctx context.Context, u *User) error {
	q := `INSERT INTO test_table (name, age) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, u.Name, u.Age); err != nil {
		return fmt.Errorf("can't save page: %w", err)
	}

	return nil
}

func (s *Storage) Remove(ctx context.Context, u *User) error {
	q := `DELETE FROM test_table WHERE name = ? AND age = ?`
	if _, err := s.db.ExecContext(ctx, q, u.Name, u.Age); err != nil {
		return fmt.Errorf("can't remove page: %w", err)
	}

	return nil
}
