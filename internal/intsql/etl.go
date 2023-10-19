package intsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log/slog"

	duckdb "github.com/marcboeker/go-duckdb"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

func GetExistingActivityIds(db *sql.DB) (strava.IntSet, error) {
	sqlText := `select activity_id from activities;`
	rows, err := db.Query(sqlText)
	if err != nil {
		return nil, fmt.Errorf("querying for existing activity ids: %w", err)
	}

	set := strava.IntSet{}
	var i int64
	for rows.Next() {
		err = rows.Scan(&i)
		if err != nil {
			return nil, fmt.Errorf("scanning existing activity ids: %w", err)
		}
		set[i] = struct{}{}
	}
	return set, nil
}

func UploadActivityJson[T util.Jsonable](db *sql.DB, activities []T) error {
	slog.Info("UploadActivityJson")

	// we use a conn as this is required for Raw.  In the future I want temp_etl to be a temp table (currently
	// temp tables cannot have extension types such as json, a bug) so it must also be handled with the conn
	conn, err := db.Conn(context.TODO())
	if err != nil {
		return fmt.Errorf("getting connection: %w", err)
	}
	defer conn.Close()

	row := QueryRowContext(context.TODO(), conn, "clear temp etl", `delete from "temp_etl";`)
	if row.Err() != nil {
		return fmt.Errorf("clear temp etl: %w", row.Err())
	}
	LogRowResponse(row, "clear temp etl")

	slog.Info("appending into temp_etl")
	err = conn.Raw(func(a any) error {
		dbConn, ok := a.(driver.Conn)
		if !ok {
			return fmt.Errorf("not a duckdb driver connection")
		}

		appender, err := duckdb.NewAppenderFromConn(dbConn, "", "temp_etl")
		if err != nil {
			return fmt.Errorf("creating appender for json etl: %w", err)
		}
		defer appender.Close()

		for _, activity := range activities {
			err = appender.AppendRow(activity.ToJson())
			if err != nil {
				return fmt.Errorf("appending row: %w", err)
			}
		}

		err = appender.Flush()
		if err != nil {
			return fmt.Errorf("flushing appender: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("conn.raw for appender: %w", err)
	}
	slog.Info("appended into temp_etl", "rows", len(activities))

	sqlText := `insert into "etl" (activity_id, json)
		select ("json"->'$.Activity.id')::int64, json("json")
		from "temp_etl";`
	row = QueryRowContext(context.TODO(), conn, "insert into etl", sqlText)
	if row.Err() != nil {
		return fmt.Errorf("insert into etl: %w", row.Err())
	}
	LogRowResponse(row, "insert into etl")

	return nil
}

func MergeActivities(db *sql.DB) error {
	// TODO
	// remove the limit for new streams

	slog.Info("MergeActivities")
	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("starting transaction: %w", err)
	}
	defer txn.Rollback()
	var dbCtx QueryRowContextable = txn

	// insert on conflict (merge) is not supported in duckdb for anything including a list type,
	// which appears to include lists nested in structs.  So we will transactionally delete all
	// rows in activities in streams with a matching activity_id and then reinsert these rows.

	// NB: we use the activities and etl tables to determine duplicates.  We must delete rows
	// from streams before activities as once we delete from activities we will no longer deem
	// the activity_ids as duplicates
	sqlText := `
		delete from streams
		where activity_id in (
			select activity_id from activity_ids_to_update);`
	row := QueryRowContext(context.TODO(), dbCtx, "delete duplicate streams", sqlText)
	if row.Err() != nil {
		return fmt.Errorf("removing duplicate activities from streams: %w", row.Err())
	}
	LogRowResponse(row, "delete duplicate streams")

	sqlText = `delete from activities
		where activity_id in (
			select activity_id from activity_ids_to_update);`
	row = QueryRowContext(context.TODO(), dbCtx, "delete duplicate activities", sqlText)
	if row.Err() != nil {
		return fmt.Errorf("removing duplicate activities from activities: %w", row.Err())
	}
	LogRowResponse(row, "delete duplicate activities")

	// insert rows, both new and as semantic updates.
	sqlText = `insert into activities select * from new_activities;`
	row = QueryRowContext(context.TODO(), dbCtx, "insert new activities", sqlText)
	if row.Err() != nil {
		return fmt.Errorf("inserting new activities: %w", row.Err())
	}
	LogRowResponse(row, "insert new activities")

	// insert the new rows into streams
	slog.Info("inserting into streams in batches")
	sqlText = `insert into streams select * from new_streams;`
	row = QueryRowContext(context.TODO(), dbCtx, "insert new streams", sqlText)
	if row.Err() != nil {
		return fmt.Errorf("inserting new streams: %w", row.Err())
	}
	LogRowResponse(row, "insert new streams")

	err = txn.Commit()
	if err != nil {
		return fmt.Errorf("committing: %w", err)
	}

	return nil
}
