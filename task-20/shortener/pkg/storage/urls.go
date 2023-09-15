package storage

import (
	"time"
	"context"
	"fmt"
	"hash/fnv"
)

type Record struct {
	Short       string    `json:"short"`
	Long  			string    `json:"long"`
	CreatedAt  	time.Time    `json:"created_at"`
}

func NewRecord(long string) Record{
	return Record{
		Short:  hash(long),
		Long: long,
		CreatedAt: time.Now(),
	}
}

func (db *DB) Add(ctx context.Context, long string) (Record, error) {
	r := NewRecord(long)
	query := `INSERT INTO urls (short_url, destination, created_at) VALUES ($1, $2, $3)`
	_, err := db.pool.Exec(ctx, query, r.Short, r.Long, r.CreatedAt)
	return r, err
}

func (db *DB) Get(ctx context.Context, short string) (*Record, error) {
	r := Record{}
	err := db.pool.QueryRow(ctx, "SELECT short, long, created_at FROM urls WHERE short = $1 LIMIT 1", short).Scan(&r.Short, &r.Long, &r.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}
