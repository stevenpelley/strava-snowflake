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
	slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "query", "row"})
	row := q.QueryRowContext(context, sqlText, args...)
	if row.Err() != nil {
		slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "error", "row"}, "error", row.Err())
	}
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
	slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "query", "rows"})
	rows, err := q.QueryContext(context, sqlText)
	if err != nil {
		slog.Info(logLabel, "sql_text", sqlText, "tag", []string{"sql", "error", "rows"}, "error", err)
	}
	return rows, err
}
