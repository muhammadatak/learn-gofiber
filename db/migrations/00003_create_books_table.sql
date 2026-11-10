-- +goose Up
CREATE TABLE books (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    isbn        VARCHAR(20)  NOT NULL UNIQUE,
    cover_id    UUID,
    created_at  TIMESTAMPTZ,
    updated_at  TIMESTAMPTZ,
    deleted_at  TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS books;
