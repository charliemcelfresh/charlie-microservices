-- migrate:up

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_users__email ON users (email);

-- migrate:down

DROP TABLE users CASCADE;
