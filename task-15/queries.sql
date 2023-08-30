-- top 3 artists by revenue (doesn't matter on their role)
SELECT DISTINCT
  people.name,
  (
    SELECT
      SUM(revenue)
    FROM
      movies
    WHERE
      movies.id IN (
        SELECT
          movie_id
        FROM
          movies_people
        WHERE
          person_id = people.id
      )
  ) AS total_revenue
FROM
  people
  JOIN movies_people ON people.id = person_id
  JOIN movies ON movies.id = movie_id
ORDER BY
  total_revenue DESC
LIMIT
  3;

-- average actor age by studio
SELECT DISTINCT
  studios.name,
  (
    SELECT
      AVG(AGE (birth_date))
    FROM
      people
    WHERE
      people.id IN (
        SELECT
          person_id
        FROM
          movies_people
        WHERE
          cast_type = 'actor'
          AND movie_id IN (
            SELECT
              movies.id
            FROM
              movies
            WHERE
              studio_id = studios.id
          )
      )
  ) AS avg_age
FROM
  studios
  JOIN movies ON studios.id = studio_id
  JOIN movies_people ON movies.id = movie_id
  JOIN people ON people.id = person_id
ORDER BY
  avg_age DESC
LIMIT
  2;

-- the latest 10 movies with a studio and director
SELECT DISTINCT
  year,
  movies.title AS title,
  studios.name AS studio,
  (
    SELECT
      name
    FROM
      people
    WHERE
      people.id in (
        SELECT
          person_id
        FROM
          movies_people
        WHERE
          cast_type = 'director'
          and movie_id = movies.id
      )
    LIMIT
      1
  ) as director
FROM
  movies
  JOIN studios ON studios.id = studio_id
  JOIN movies_people ON movies.id = movie_id
  JOIN people ON people.id = person_id
ORDER BY
  year DESC
limit
  10;