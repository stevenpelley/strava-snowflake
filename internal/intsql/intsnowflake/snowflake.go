package intsnowflake

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/snowflakedb/gosnowflake"
	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stevenpelley/strava-snowflake/internal/util"
)

type SFStrava struct {
	db             *sql.DB
	configFileName string
	etlTableName   string
}

// prove it is a StravaDatabase
var _ intsql.StravaDatabase = &SFStrava{}

func New(configFileName string, etlTableName string) SFStrava {
	return SFStrava{configFileName: configFileName, etlTableName: etlTableName}
}

// OpenDB implements intsql.StravaDatabase.
func (sdb *SFStrava) OpenDB() error {
	config, err := ToSnowflakeConfig(sdb.configFileName)
	if err != nil {
		return err
	}

	dsn, err := gosnowflake.DSN(config)
	if err != nil {
		return fmt.Errorf("gosnowflake.DSN: %w", err)
	}

	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open snowflake: %w", err)
	}

	sdb.db = db
	return nil
}

func (sdb *SFStrava) DB() *sql.DB {
	return sdb.db
}

// InitAndValidateSchema implements intsql.StravaDatabase.
func (sdb *SFStrava) InitAndValidateSchema() error {
	// just not going to do this here
	return nil
}

// FilterKnownActivityIds implements intsql.StravaDatabase.
func (sdb *SFStrava) FilterKnownActivityIds(activityIds []int64) ([]int64, error) {
	conn, err := sdb.db.Conn(context.Background())
	if err != nil {
		return nil, fmt.Errorf("conn: %w", err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(context.Background(), "create temp table activity_ids (n number);")
	if err != nil {
		return nil, fmt.Errorf("create temp table: %w", err)
	}

	_, err = conn.ExecContext(
		context.Background(), "insert into activity_ids values (?);", activityIds)
	if err != nil {
		return nil, fmt.Errorf("insert activity ids: %w", err)
	}

	rows, err := conn.QueryContext(
		context.Background(),
		"select n from activity_ids where not in (select activity:id::number from etl);")
	if err != nil {
		return nil, fmt.Errorf("select: %w", err)
	}

	result := make([]int64, 0)
	for rows.Next() {
		var i int64
		err = rows.Scan(&i)
		if err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
		result = append(result, i)
	}

	return result, nil
}

// UploadActivityJson implements intsql.StravaDatabase.
func (sdb *SFStrava) UploadActivityJson(activities []util.Jsonable) error {
	s := make([]string, len(activities))
	for idx, activity := range activities {
		s[idx] = activity.ToJson()
	}
	_, err := sdb.db.ExecContext(
		context.Background(),
		fmt.Sprintf(
			`insert into %v (data) select parse_json($1) from values (?);`,
			sdb.etlTableName),
		gosnowflake.Array(s))
	if err != nil {
		return fmt.Errorf("insert: %w", err)
	}

	return nil
}

// MergeActivities implements intsql.StravaDatabase.
func (sdb *SFStrava) MergeActivities() error {
	// no merge necessary in snowflake
	return nil
}

// Close implements intsql.StravaDatabase.
func (sdb *SFStrava) Close() error {
	if sdb.db == nil {
		return nil
	}
	return sdb.db.Close()
}
