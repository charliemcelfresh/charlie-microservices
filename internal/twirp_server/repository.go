package twirp_server

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type database interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Transaction struct {
	ID            string `db:"id"`
	UserAccountID string `db:"user_account_id"`
	CardNetwork   string `db:"card_network"`
	MerchantID    string `db:"merchant_id"`
	MerchantName  string `db:"merchant_name"`
	Total         string `db:"total"`
	CreatedAt     string `db:"created_at"`
}

type repository struct {
	db database
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db,
	}
}

func (r repository) GetTransactions(ctx context.Context, userID, cardNetwork string) ([]Transaction, error) {
	transactions := []Transaction{}
	statement := `
		SELECT
			t.id,
			t.user_account_id,
			t.card_network,
			t.merchant_id,
			t.merchant_name,
			t.total::TEXT,
			t.created_at
		FROM
		    transactions t
		JOIN
			user_accounts ua ON ua.id = t.user_account_id
		WHERE
			ua.user_id = $1 AND
			ua.card_network = $2
		ORDER BY
		    t.created_at DESC
    `
	err := r.db.SelectContext(ctx, &transactions, statement, userID, cardNetwork)

	return transactions, err
}
