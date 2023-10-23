package intduckdb

import (
	"os"
	"syscall"
	"testing"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stretchr/testify/require"
)

func deleteTestFile(t *testing.T, testFile string) {
	if testFile == "" {
		return
	}

	err := os.Remove(testFile)
	if err != nil {
		require.ErrorIs(t, err, syscall.ENOENT)
	}
	err = os.Remove(testFile + ".wal")
	if err != nil {
		require.ErrorIs(t, err, syscall.ENOENT)
	}
}

type closeDb int64

const (
	Invalid closeDb = iota
	Close
	LeaveOpen
)

func DBTest(t *testing.T, testFile string, closeDb closeDb, f func(sdb *DuckdbStrava)) {
	deleteTestFile(t, testFile)
	if closeDb == Close {
		defer deleteTestFile(t, testFile)
	}
	sdb := New(testFile)
	err := sdb.OpenDB()
	require.NoError(t, err)
	defer sdb.Close()

	require.NoError(t, sdb.InitAndValidateSchema())
	f(&sdb)
}

func TestUploadActivityJson(t *testing.T) {
	testFile := ""
	DBTest(t, testFile, Close, func(sdb *DuckdbStrava) {
		intsql.HelperTestUploadActivityJson(t, sdb)
	})
}

func TestMergeActivities(t *testing.T) {
	testFile := ""
	DBTest(t, testFile, Close, func(sdb *DuckdbStrava) {
		intsql.HelperTestMergeActivities(t, sdb)
	})
}

// this is testing a common utility but needs a database to do so.  Duckdb is simplest and
// most appropriate for running tests.
// this is testing a common utility but needs a database to do so.  Duckdb is simplest and
// most appropriate for running tests.
func TestScanColumns(t *testing.T) {
	require := require.New(t)
	DBTest(t, "", Close, func(sdb *DuckdbStrava) {
		row := sdb.db.QueryRow("create table t(c1 int64, c2 string);")
		require.NoError(row.Err())

		row = sdb.db.QueryRow("insert into t values (1, 'asdf'), (2, 'qwer');")
		require.NoError(row.Err())

		rows, err := sdb.db.Query("select * from t;")
		require.NoError(err)
		defer rows.Close()

		var ints []int64
		var strings []string
		err = intsql.ScanColumns(rows, &ints, &strings)
		require.NoError(err)

		require.EqualValues(ints, []int64{1, 2})
		require.EqualValues(strings, []string{"asdf", "qwer"})

		rows, err = sdb.db.Query("select * from t;")
		require.NoError(err)
		defer rows.Close()
		intsql.ScanColumnsAndCompare(t, rows, []int64{1, 2}, []string{"asdf", "qwer"})
	})
}
