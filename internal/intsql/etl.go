package intsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	duckdb "github.com/marcboeker/go-duckdb"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

func GetExistingActivityIds(db *sql.DB) (strava.IntSet, error) {
	sqlText := fmt.Sprintf(`select activity_id from %v;`, activitiesTable)
	rows, err := db.Query(sqlText)
	if err != nil {
		return nil, fmt.Errorf("error querying for existing activity ids: %w", err)
	}

	set := strava.IntSet{}
	var i int64
	for rows.Next() {
		err = rows.Scan(&i)
		if err != nil {
			return nil, fmt.Errorf("error scanning existing activity ids: %w", err)
		}
		set[i] = struct{}{}
	}
	return set, nil
}

func UploadActivityJson[T util.Jsonable](db *sql.DB, activities []T) error {
	conn, err := db.Conn(context.TODO())
	if err != nil {
		return fmt.Errorf("error getting connection: %w", err)
	}
	defer conn.Close()

	rows, err := conn.QueryContext(
		context.TODO(),
		fmt.Sprintf("delete from \"%v\";", tempEtlTable))
	if err != nil {
		return fmt.Errorf("error clearing temp etl table: %w", err)
	}
	defer rows.Close()

	err = conn.Raw(func(a any) error {
		dbConn, ok := a.(driver.Conn)
		if !ok {
			return fmt.Errorf("not a duckdb driver connection")
		}

		appender, err := duckdb.NewAppenderFromConn(dbConn, "", tempEtlTable)
		if err != nil {
			return fmt.Errorf("error creating appender for json etl: %w", err)
		}
		defer appender.Close()

		for _, activity := range activities {
			err = appender.AppendRow(activity.ToJson())
			if err != nil {
				return fmt.Errorf("error appending row: %w", err)
			}
		}

		err = appender.Flush()
		if err != nil {
			return fmt.Errorf("error flushing appender: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error running conn.raw for appender: %w", err)
	}

	_, err = conn.ExecContext(
		context.TODO(),
		fmt.Sprintf("insert into \"%v\" (activity_id, json) select (\"json\"->'$.Activity.id')::int64, json(\"json\") from \"%v\";",
			etlTable, tempEtlTable))
	if err != nil {
		return fmt.Errorf("error copying from temp_etl to etl: %w", err)
	}

	return nil
}

func MergeActivities(db *sql.DB) error {
	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer txn.Rollback()

	// insert on conflict (merge) is not supported in duckdb for anything including a list type,
	// which appears to include lists nested in structs.  So we will transactionally delete all
	// rows in activities in streams with a matching activity_id and then reinsert these rows.

	// NB: we use the activities and etl tables to determine duplicates.  We must delete rows
	// from streams before activities as once we delete from activities we will no longer deem
	// the activity_ids as duplicates
	row := txn.QueryRowContext(context.TODO(), `
		delete from streams
		where activity_id in (
			select activity_id from activity_ids_to_update);`)
	if row.Err() != nil {
		return fmt.Errorf("removing duplicate activities from streams: %w", row.Err())
	}

	row = txn.QueryRowContext(context.TODO(), `
		delete from activities
		where activity_id in (
			select activity_id from activity_ids_to_update);`)
	if row.Err() != nil {
		return fmt.Errorf("removing duplicate activities from activities: %w", row.Err())
	}

	// insert rows, both new and as semantic updates.
	row = txn.QueryRowContext(context.TODO(), `insert into activities select * from new_activties;`)
	if row.Err() != nil {
		return fmt.Errorf("inserting new activities: %w", row.Err())
	}

	// insert the new rows into streams
	row = txn.QueryRowContext(context.TODO(), `insert into streams select * from new_streams;`)
	if row.Err() != nil {
		return fmt.Errorf("inserting new activities: %w", row.Err())
	}

	err = txn.Commit()
	if err != nil {
		return fmt.Errorf("error committing: %w", err)
	}

	return nil
}
