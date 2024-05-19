-- migrate:up

CREATE TABLE merchants (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_merchants__name ON merchants (name);

-- migrate:down

DROP TABLE merchants CASCADE;
