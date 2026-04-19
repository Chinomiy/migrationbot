-- +goose Up
ALTER TABLE country ADD COLUMN published BOOLEAN;
ALTER TABLE country ALTER COLUMN published SET NOT NULL;

-- +goose Down
ALTER TABLE country DROP COLUMN published;