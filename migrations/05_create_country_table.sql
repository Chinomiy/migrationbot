-- +goose Up
CREATE TABLE IF NOT EXISTS country
(
    id         BIGSERIAL PRIMARY KEY,
    code       TEXT        NOT NULL,
    name TEXT NOT NULL,
    description   TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS trip_type (

    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    callback TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS country_trip_type (
    country_id INTEGER REFERENCES country(id) ,
    trip_type_id INTEGER REFERENCES trip_type(id) ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (country_id, trip_type_id)
);

CREATE TABLE IF NOT EXISTS country_trip_content (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    country_id INTEGER REFERENCES country(id) ,
    trip_type_id INTEGER REFERENCES trip_type(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (country_id, trip_type_id)
);
-- +goose Down
DROP TABLE country_trip_content;
DROP TABLE country_trip_type;
DROP TABLE trip_type;
DROP TABLE country;