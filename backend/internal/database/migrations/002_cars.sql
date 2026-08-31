CREATE TABLE IF NOT EXISTS cars (
    id           BIGSERIAL PRIMARY KEY,
    make_id      BIGINT NOT NULL REFERENCES makes (id),
    model        TEXT NOT NULL,
    year         INT NOT NULL,
    price        BIGINT NOT NULL,
    mileage      BIGINT NOT NULL,
    description  TEXT NOT NULL DEFAULT '',
    seller_name  TEXT NOT NULL,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS cars_make_id_idx ON cars (make_id);
CREATE INDEX IF NOT EXISTS cars_year_idx ON cars (year);
CREATE INDEX IF NOT EXISTS cars_price_idx ON cars (price);
