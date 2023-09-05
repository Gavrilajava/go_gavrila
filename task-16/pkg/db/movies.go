package db

import (
	"context"
	// "fmt"

	"github.com/jackc/pgx/v5"
)



type Interface interface {
	Movies(ctx context.Context, studioId int) ([]Movie, error)
	AddMovies(ctx context.Context, books []Movie) error
	UpdateMovie(ctx context.Context, books []Movie) error
	DeleteMovie(ctx context.Context, books []Movie) error
}

type Movie struct {
	ID          int8
	Title       string
	Year        int
	Revenue     int
	StidioID    int
	Guidance 		string
}


func (db *DB) Movies(ctx context.Context, studioId int) ([]Movie, error) {

	rows, err := db.pool.Query(ctx, "SELECT id, title, year, revenue, guidance, studio_id FROM movies WHERE $1 = 0 OR studio_id = $1", studioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := make([]Movie, 0)

	for rows.Next() {
		m := &Movie{}
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Year,
			&m.Revenue,
			&m.Guidance,
			&m.StidioID,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, *m)
	}

	return movies, nil
}

func (db *DB) AddMovies(ctx context.Context, movies []Movie) error {
	batch := &pgx.Batch{}
	for _, m := range movies {
		batch.Queue(
			"INSERT INTO movies (title, year, revenue, guidance, studio_id) VALUES ($1, $2, $3, $4, $5)",
			m.Title, m.Year, m.Revenue, m.Guidance, m.StidioID,
		)
	}

	return db.pool.SendBatch(ctx, batch).Close()
}

func (db *DB) UpdateMovie(ctx context.Context, m Movie) (err error) {
	_, err = db.pool.Exec(ctx,
		"UPDATE movies SET title = $1, year = $2, revenue = $3, guidance = $4, studio_id = $5 WHERE id = $6",
		m.Title, m.Year, m.Revenue, m.Guidance, m.StidioID, m.ID,
	)
	return err
}

func (db *DB) DeleteMovie(ctx context.Context, id int) (err error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM movies_people WHERE movie_id = $1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM movies WHERE id = $1", id)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

