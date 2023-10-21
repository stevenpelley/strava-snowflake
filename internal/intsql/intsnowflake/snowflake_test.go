package intsnowflake

import (
	"database/sql"
	"testing"

	"github.com/snowflakedb/gosnowflake"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	require := require.New(t)
	sfConfig, err := ToSnowflakeConfig("snowflake_config.json")
	require.NoError(err)

	dsn, err := gosnowflake.DSN(sfConfig)
	require.NoError(err)
	db, err := sql.Open("snowflake", dsn)
	require.NoError(err)
	defer db.Close()

	row := db.QueryRow("select 1;")
	require.NoError(row.Err())
	var i int64
	err = row.Scan(&i)
	require.NoError(err)
	require.Equal(int64(1), i)
}
