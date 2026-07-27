CREATE TABLE IF NOT EXISTS makes (
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS makes_name_lower_idx ON makes (LOWER(name));

INSERT INTO makes (name)
SELECT 'Toyota'
WHERE NOT EXISTS (SELECT 1 FROM makes WHERE LOWER(name) = 'toyota');

INSERT INTO makes (name)
SELECT 'Volkswagen'
WHERE NOT EXISTS (SELECT 1 FROM makes WHERE LOWER(name) = 'volkswagen');
