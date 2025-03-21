package write

import (
	"context"
	"database/sql"

	iwrite "shared-wallet-service/infrastructure/database/interfaces/write"
)

type write struct {
	db *sql.DB
}

func NewSqlWrite(db *sql.DB) iwrite.IWrite {
	return &write{
		db: db,
	}
}

func (w *write) ExecuteContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return w.db.ExecContext(ctx, query, args...)
}

func (w *write) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return w.db.BeginTx(ctx, opts)
}

func (w *write) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (w *write) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
