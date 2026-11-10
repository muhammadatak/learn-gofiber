-- +goose Up
CREATE TABLE journals (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    book_id     UUID        NOT NULL REFERENCES books (id),
    stock_code  VARCHAR(50) NOT NULL REFERENCES book_stocks (code),
    customer_id UUID        NOT NULL REFERENCES customers (id),
    status      VARCHAR(50) NOT NULL,
    borrowed_at TIMESTAMPTZ,
    returned_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS journals;
