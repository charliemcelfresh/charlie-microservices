-- migrate:up

    CREATE INDEX idx_transactions_card_network ON transactions (card_network);

-- migrate:down

    DROP INDEX idx_transactions_card_network;