package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	driver          = "postgres"
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)

type Storage struct {
	db *sql.DB
}

func New(datasource string) (*Storage, error) {
	db, err := sql.Open(driver, datasource)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	log.Printf("connected to %s database", driver)

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) SaveScore(ctx context.Context, userId int, time time.Duration) (int64, error) {
	var id int64

	query := "INSERT INTO scores (user_id, time) VALUES ($1, $2) RETURNING id"

	err := s.db.QueryRowContext(ctx, query, userId, time).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to save score: %w", err)
	}

	return id, nil
}
