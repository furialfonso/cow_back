package read

import (
	"context"
	"database/sql"

	iread "shared-wallet-service/infrastructure/database/interfaces/read"
)

type read struct {
	db *sql.DB
}

func NewSqlRead(db *sql.DB) iread.IRead {
	return &read{
		db: db,
	}
}

func (r *read) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return r.db.QueryContext(ctx, query, args...)
}

func (r *read) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return r.db.QueryRowContext(ctx, query, args...)
}
