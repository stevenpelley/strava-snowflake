package intsnowflake

import (
	"testing"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stretchr/testify/require"
)

func DBTest(t *testing.T, f func(sdb *SFStrava)) {
	sdb := New("../../../snowflake_config.json")
	err := sdb.OpenDB()
	require.NoError(t, err)
	defer sdb.Close()

	require.NoError(t, sdb.InitAndValidateSchema())
	f(&sdb)
}

func TestConnect(t *testing.T) {
	DBTest(t, func(sdb *SFStrava) {
		require := require.New(t)
		row := sdb.db.QueryRow("select 1;")
		require.NoError(row.Err())
		var i int64
		err := row.Scan(&i)
		require.NoError(err)
		require.Equal(int64(1), i)
	})
}

func TestUploadActivityJson(t *testing.T) {
	DBTest(t, func(sdb *SFStrava) {
		intsql.HelperTestUploadActivityJson(t, sdb)
	})
}

func TestMergeActivities(t *testing.T) {
	//testFile := ""
	//DBTest(t, testFile, Close, func(sdb *DuckdbStrava) {
	//	intsql.HelperTestMergeActivities(t, sdb)
	//})
}
