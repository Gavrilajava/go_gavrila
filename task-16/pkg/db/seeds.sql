
INSERT INTO studios (name) VALUES ('Disney'),  ('Lenfilm');

INSERT INTO people (name, birth_date) VALUES 
	('Quentin Tarantino', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Nikita Mikhalkov', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Brad Pitt', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Leonardo Di Caprio', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Samuel L Jackson', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Ivan Urgant', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Semen Slepakov', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) ),
	('Alexander Tsekalo', (SELECT RANDOM() * ('1960-01-01'::timestamp - '1990-12-31'::timestamp) + '1990-12-31'::timestamp) );

INSERT INTO movies (title, year, revenue, guidance, studio_id) VALUES 
	('Yolki', 2010, 1000000, 'PG-10', 1),
	('Yolki 2', 2011, 2000000, 'PG-13', 2),
	('Yolki 3', 2012, 3000000, 'PG-18', 1),
	('Kill Bill', 2013, 4000000, 'PG-10', 2),
	('Kill Bill 2', 2014, 5000000, 'PG-13', 1),
	('Kill Bill 3', 2015, 6000000, 'PG-18', 2);

DO
$$
DECLARE
    m integer;
    p integer;
BEGIN
    FOR m in SELECT id FROM movies 
    LOOP 
      INSERT INTO movies_people (movie_id, cast_type, person_id ) VALUES (m, 'director', (SELECT id FROM people ORDER BY RANDOM() LIMIT 1));
      FOR p in SELECT id FROM people ORDER BY RANDOM() LIMIT RANDOM() * 5
      LOOP
        INSERT INTO movies_people (movie_id, cast_type, person_id ) VALUES (m, 'actor', p);
      END LOOP;
    END LOOP;
END;
$$;