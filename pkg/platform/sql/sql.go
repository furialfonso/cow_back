package sql

import (
	"context"
	"database/sql"
)

type read struct {
	db *sql.DB
}

type write struct {
	db *sql.DB
}

func NewSqlRead(db *sql.DB) IRead {
	return &read{
		db: db,
	}
}

func NewSqlWrite(db *sql.DB) IWrite {
	return &write{
		db: db,
	}
}

func (r *read) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return r.db.QueryContext(ctx, query, args...)
}

func (r *read) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return r.db.QueryRowContext(ctx, query, args...)
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
