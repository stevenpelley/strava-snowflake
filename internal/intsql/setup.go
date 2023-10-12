package intsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"flag"
	"log/slog"

	duckdb "github.com/marcboeker/go-duckdb"
)

type DuckdbFlags struct {
	DbFileName      string
	StreamsEtlLimit int64
}

func (df *DuckdbFlags) InitFlags() {
	flag.StringVar(&df.DbFileName, "duckdbfile", "", "duckdb database file (empty for memory database)")
	flag.Int64Var(&df.StreamsEtlLimit, "streamsetllimit", 10000, "limit number of streams rows to ingest at a time to control memory usage")
}

func (df *DuckdbFlags) PostProcessFlags() {
}

func OpenDB(dbFileName string) (*sql.DB, error) {
	slog.Info("opening database", "filename", dbFileName)
	connector, err := duckdb.NewConnector(dbFileName, func(execer driver.ExecerContext) error {
		bootQueries := []string{
			"INSTALL 'json'",
			"LOAD 'json'",
		}

		for _, qry := range bootQueries {
			_, err := execer.ExecContext(context.TODO(), qry, nil)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	db := sql.OpenDB(connector)
	return db, nil
}
