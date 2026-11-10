-- +goose Up
ALTER TABLE journals ADD COLUMN due_at TIMESTAMPTZ;

-- +goose Down
ALTER TABLE journals DROP COLUMN IF EXISTS due_at;
