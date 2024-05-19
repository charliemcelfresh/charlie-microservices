-- migrate:up

CREATE TABLE transactions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_account_id uuid NOT NULL references user_accounts(id),
    card_network cardnetwork_enum NOT NULL,
    merchant_id TEXT NOT NULL,
    merchant_name TEXT NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    created_at timestamptz DEFAULT NOW(),
    updated_at timestamptz DEFAULT NOW()
);

CREATE INDEX idx_transactions__user_account_id ON transactions (user_account_id);
CREATE INDEX idx_transactions__merchant_id ON transactions (merchant_id);
CREATE INDEX idx_transactions__created_at ON transactions (created_at);

-- migrate:down

DROP TABLE transactions CASCADE;
