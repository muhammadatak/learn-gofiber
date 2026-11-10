-- +goose Up
CREATE TABLE charges (
    id             UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    journal_id     UUID          NOT NULL REFERENCES journals (id),
    days_late      INT           NOT NULL,
    daily_late_fee NUMERIC(12, 2) NOT NULL,
    total          NUMERIC(12, 2) NOT NULL,
    user_id        UUID          NOT NULL REFERENCES users (id),
    created_at     TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS charges;
