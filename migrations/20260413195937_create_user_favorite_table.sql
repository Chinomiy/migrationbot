-- +goose Up
-- USER.ID + country.ID + trip_type.ID есть в country_trip_type
CREATE TABLE user_favorite(
    user_id BIGINT REFERENCES users(tg_id) ON DELETE CASCADE,
    country_trip_content_id INT REFERENCES country_trip_content(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, country_trip_content_id)
);


-- +goose Down
SELECT 'down SQL query';
