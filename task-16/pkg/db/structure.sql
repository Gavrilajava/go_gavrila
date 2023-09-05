

DROP TABLE IF EXISTS people CASCADE;
CREATE TABLE people (
    id SERIAL PRIMARY KEY,
    name varchar(200) NOT NULL DEFAULT '',
    birth_date date NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(name)
);

DROP TYPE IF EXISTS guidances CASCADE;
CREATE TYPE guidances AS ENUM ('PG-10', 'PG-13', 'PG-18');

DROP TABLE IF EXISTS movies CASCADE;
CREATE TABLE movies (
    id BIGSERIAL PRIMARY KEY,
    title varchar(200) NOT NULL,
    year INTEGER NOT NULL DEFAULT 1800,
    revenue INTEGER DEFAULT 0,
    guidance guidances NOT NULL,
    studio_id INTEGER REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(title, year),
    CONSTRAINT year_not_less_1800 check (year >= 1800)
);

DROP TYPE IF EXISTS cast_types CASCADE;
CREATE TYPE cast_types AS ENUM ('actor', 'director');

DROP TABLE IF EXISTS movies_people;
CREATE TABLE movies_people (
    id SERIAL PRIMARY KEY,
    movie_id INTEGER NOT NULL REFERENCES movies(id),
    person_id INTEGER NOT NULL REFERENCES people(id),
    cast_type cast_types NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE(movie_id, person_id, cast_type)
);

DROP TABLE IF EXISTS studios CASCADE;
CREATE TABLE studios (
  id SERIAL PRIMARY KEY,
  name varchar(200) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  UNIQUE(name)
);