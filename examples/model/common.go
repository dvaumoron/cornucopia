// Generated from ../test.crn - do not edit.

package model

import (
	"context"
	"database/sql"
	"time"
)

const timeout = time.Duration(int64(5)) * time.Second

type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type RowQueryerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}
