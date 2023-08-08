/*
    В данной схеме создано 4 таблицы: `studios`, `actors`, `directors` и `movies`.
    Таблица `studios` содержит информацию о студиях, таблицы `actors` и `directors` содержат информацию об актерах и режиссёрах соответственно.
    Таблица `movies` хранит информацию о фильмах, включая название, год выхода, сборы и рейтинг.
    В таблице `movies` есть столбец `studio_id`, который ссылается на таблицу `studios`, указывая принадлежность фильма к конкретной студии.
    Также созданы таблицы `movie_actors` и `movie_directors`, которые связывают фильмы с соответствующими актерами и режиссёрами.
    С использованием `CONSTRAINT unique_movie_name_per_year UNIQUE (name, year_of_release)` добавлена проверка на уникальность названия фильма в определенный год.
*/

CREATE TABLE studios
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE actors
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    date_of_birth DATE         NOT NULL
);

CREATE TABLE directors
(
    id            SERIAL PRIMARY KEY,
    name          VARCHAR(100) NOT NULL,
    date_of_birth DATE         NOT NULL
);

CREATE TABLE movies
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100)   NOT NULL,
    year_of_release INTEGER        NOT NULL CHECK (year_of_release >= 1800),
    box_office      DECIMAL(10, 2) NOT NULL,
    rating          VARCHAR(10)    NOT NULL CHECK (rating IN ('PG-10', 'PG-13', 'PG-18')),
    studio_id       INTEGER        NOT NULL REFERENCES studios (id),
    CONSTRAINT unique_movie_name_per_year UNIQUE (name, year_of_release)
);

CREATE TABLE movie_actors
(
    movie_id INTEGER NOT NULL REFERENCES movies (id),
    actor_id INTEGER NOT NULL REFERENCES actors (id),
    PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE movie_directors
(
    movie_id    INTEGER NOT NULL REFERENCES movies (id),
    director_id INTEGER NOT NULL REFERENCES directors (id),
    PRIMARY KEY (movie_id, director_id)
);