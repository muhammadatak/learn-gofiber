-- +goose Up
CREATE TABLE media (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    path       TEXT NOT NULL,
    created_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS media;
