-- +goose Up
    ALTER TABLE country_trip_content ADD COLUMN requirements TEXT;
    ALTER TABLE country_trip_content ALTER COLUMN requirements SET NOT NULL;

-- +goose Down
    ALTER TABLE country_trip_content DROP COLUMN requirements;