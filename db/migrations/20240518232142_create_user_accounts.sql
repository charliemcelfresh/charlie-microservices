-- migrate:up

CREATE TYPE cardnetwork_enum AS ENUM ('MasterCard', 'Amex', 'Visa');


CREATE TABLE user_accounts (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL references users(id),
    card_network cardnetwork_enum NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE INDEX idx_user_accounts__user_id ON user_accounts (user_id);
CREATE INDEX idx_user_accounts__card_network ON user_accounts (card_network);

-- migrate:down

DROP TYPE cardnetwork_enum CASCADE;
DROP TABLE user_accounts;
