-- +goose Up
ALTER TABLE users ADD COLUMN tg_username TEXT;
ALTER TABLE users ALTER COLUMN tg_username SET NOT NULL;

-- +goose Down
ALTER TABLE users DROP COLUMN tg_username;