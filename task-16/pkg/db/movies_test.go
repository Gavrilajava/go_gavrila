package db

import (
	"context"
	"log"
	"os"
	"testing"
	"reflect"
)

var testDB *DB


func TestMain(m *testing.M) {

	var err error

	testDB, err = New("localhost",  5432, "lmdb", "georgegavrilchik", "", "disable", 10)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer testDB.pool.Close()


	schema, err := os.ReadFile("structure.sql")
	if err != nil {
		log.Fatalf("Could not read schema: %s", err)
	}

	_, err = testDB.pool.Exec(context.Background(), string(schema))
	if err != nil {
		log.Fatalf("Could not apply schema: %s", err)
	}

	seeds, err := os.ReadFile("seeds.sql")
	if err != nil {
		log.Fatalf("Could not read data sql: %s", err)
	}

	_, err = testDB.pool.Exec(context.Background(), string(seeds))
	if err != nil {
		log.Fatalf("Could not insert data: %s", err)
	}

	m.Run()

}

func TestDB_Movies(t *testing.T) {
	ctx := context.Background()

	got, err := testDB.Movies(ctx, 1)
	want := []Movie{
		{
			ID:           1,
			Title:        "Yolki",
			Year:       	2010,
			Guidance:     "PG-10",
			Revenue:      1000000,
			StidioID:    	1,
		},
		{
			ID:           3,
			Title:        "Yolki 3",
			Year:       	2012,
			Guidance:     "PG-18",
			Revenue:      3000000,
			StidioID:    	1,
		},
		{
			ID:           5,
			Title:        "Kill Bill 2",
			Year:       	2014,
			Guidance:     "PG-13",
			Revenue:      5000000,
			StidioID:    	1,
		},

	}

	if err != nil {
		t.Errorf("TestDB_Movies error = %v, wantErr %v", err, true)
		return
	}
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("TestDB_Movies got %v, want %v", got, want)
	}
}

func TestDB_AddMovies(t *testing.T) {
	ctx := context.Background()
	movies := []Movie{
		{
			Title:        "Yolki IV",
			Year:       	2023,
			Guidance:     "PG-18",
			Revenue:      1200000,
			StidioID:    	1,
		},
	}
	err := testDB.AddMovies(ctx, movies)
	if err != nil {
		t.Errorf("TestDB_AddMovies error = %v", err)
		return
	}
	m, err := testDB.Movies(ctx, 0)
	if err != nil {
		t.Errorf("TestDB_AddMovies error = %v", err)
		return
	}
	got := m[len(m)-1]
	want := movies[len(movies)-1]
	got.ID = want.ID
	if got != want {
		t.Errorf("TestDB_AddMovies got = %v, want %v", got, want)
		return
	}

}

func TestDB_UpdateMovie(t *testing.T) {
	ctx := context.Background()
	want := Movie{
		ID:          1,	
		Title:        "Yolki V",
		Year:       	2023,
		Guidance:     "PG-18",
		Revenue:      1200000,
		StidioID:    	5,
	}

	err := testDB.UpdateMovie(ctx, want)
	if err != nil {
		t.Errorf("TestDB_AddMovies error = %v", err)
		return
	}
	m, err := testDB.Movies(ctx, 5)
	if err != nil {
		t.Errorf("TestDB_AddMovies error = %v", err)
		return
	}
	got := m[0]
	if got != want {
		t.Errorf("TestDB_UpdateMovie got = %v, want %v", got, want)
		return
	}

}

func TestDB_DeleteMovie(t *testing.T) {
	ctx := context.Background()

	want := []Movie{

		{
			ID:           4,
			Title:        "Kill Bill",
			Year:       	2013,
			Guidance:     "PG-10",
			Revenue:      4000000,
			StidioID:    	2,
		},
		{
			ID:           6,
			Title:        "Kill Bill 3",
			Year:       	2015,
			Guidance:     "PG-18",
			Revenue:      6000000,
			StidioID:    	2,
		},

	}

	err := testDB.DeleteMovie(ctx, 2)
	if err != nil {
		t.Errorf("TestDB_DeleteMovie error = %v", err)
		return
	}
	got, err := testDB.Movies(ctx, 2)
	if err != nil {
		t.Errorf("TestDB_DeleteMovie error = %v", err)
		return
	}

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("TestDB_DeleteMovie got %v, want %v", got, want)
	}

}



// func TestDB_DeleteFilms(t *testing.T) {
// 	ctx := context.Background()
// 	err := testDB.DeleteFilm(ctx, 1)
// 	if err != nil {
// 		t.Errorf("DB.Films() error = %v, wantErr %v", err, true)
// 		return
// 	}

// }
