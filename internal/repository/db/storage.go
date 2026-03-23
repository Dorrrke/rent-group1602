package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) (*Storage, error) {
	conn, err := pgxpool.New(
		context.TODO(),
		"TODO",
	)
	if err != nil {
		return nil, err
	}
	return &Storage{
		conn: conn,
	}, nil
}

func (s *Storage) Close() {
	s.conn.Close()
}
