-- +goose Up
CREATE TABLE book_stocks (
    book_id     UUID        NOT NULL REFERENCES books (id),
    code        VARCHAR(50) PRIMARY KEY,
    status      VARCHAR(50) NOT NULL DEFAULT 'AVAILABLE',
    borrower_id UUID REFERENCES customers (id),
    borrowed_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS book_stocks;
