package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

func New(host string, port int, dbname string, user string, password string, ssl string, max_conns int) (*DB, error) {
	ctx := context.Background()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s  sslmode=%s pool_max_conns=%d",
		host, port, dbname, user, password, ssl, max_conns,
	)
	pool, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		return nil, err
	}

	schema, err := os.ReadFile("structure.sql")
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(context.Background(), string(schema))
	if err != nil {
		return nil, err
	}
	return &DB{pool}, nil
}


