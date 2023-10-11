package intsql

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type QueryRowContextable interface {
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

func QueryRowContext(context context.Context, q QueryRowContextable, logLabel string, sqlText string, args ...any) *sql.Row {
	row := q.QueryRowContext(context, sqlText, args...)
	slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "query", "row"}, "err", row.Err())
	return row
}

func LogRowResponse(row *sql.Row, logLabel string) error {
	var s string
	err := row.Scan(&s)
	if err != nil {
		return fmt.Errorf("%v on scan: %w", logLabel, err)
	}
	slog.Info(logLabel,
		"tag", []string{"sql", "response"},
		"response", s)
	return nil
}

type QueryContextable interface {
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

func QueryContext(context context.Context, q QueryContextable, logLabel string, sqlText string) (*sql.Rows, error) {
	rows, err := q.QueryContext(context, sqlText)
	slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "query", "rows"}, "err", err)
	return rows, err
}
