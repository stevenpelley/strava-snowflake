package intsql

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type StringJsonable string

func (sj StringJsonable) ToJson() string {
	return string(sj)
}

func deleteTestFile(t *testing.T, testFile string) {
	files, err := filepath.Glob(testFile + "*")
	require.NoError(t, err)
	for _, f := range files {
		err := os.Remove(f)
		require.NoError(t, err)
	}
}

func DBTest(t *testing.T, testFile string, f func(db *sql.DB)) {
	deleteTestFile(t, testFile)
	//defer deleteTestFile(testFile)
	db, err := OpenDB(testFile)
	require.NoError(t, err)
	defer db.Close()

	require.NoError(t, InitAndValidateSchema(db))
	f(db)
}

func ScanColumns(rows *sql.Rows, slicesOut ...any) error {
	numCols := len(slicesOut)
	sliceTypes := make([]reflect.Type, numCols)
	scanArgs := make([]interface{}, numCols)

	for i, s := range slicesOut {
		t := reflect.TypeOf(s)
		if t.Kind() != reflect.Pointer {
			return fmt.Errorf("ScanColumns col %v not a pointer", i)
		}

		elem1 := t.Elem()
		if elem1.Kind() != reflect.Slice {
			return fmt.Errorf("ScanColumns col %v not a slice pointer", i)
		}

		elem2 := elem1.Elem()
		sliceTypes[i] = elem2
		thisValue := reflect.New(elem2)
		pointerValue := thisValue.Addr()
		scanArgs[i] = pointerValue.Interface()
	}

	for rows.Next() {
		if rows.Err() != nil {
			return rows.Err()
		}

		err := rows.Scan(scanArgs...)
		if err != nil {
			return err
		}

		for i := 0; i < numCols; i++ {
			// TODO keep going here.
			// append the value to the slice.
		}
	}

	return nil
}

type intoSliceScanner struct {
	vals []interface{}
}

func (scanner *intoSliceScanner) Scan(src any) error {
	var obj interface{}
	switch v := src.(type) {
	case []byte:
		c := make([]byte, len(v))
		copy(c, v)
		obj = c
	default:
		obj = v
	}
	scanner.vals = append(scanner.vals, obj)
	return nil
}

func ScanAll(t *testing.T, rows *sql.Rows) [][]interface{} {
	columns, err := rows.Columns()
	require.NoError(t, err)
	numCols := len(columns)
	scanners := make([]intoSliceScanner, numCols)
	anys := make([]any, numCols)
	for i := 0; i < numCols; i++ {
		anys[i] = &scanners[i]
	}

	for rows.Next() {
		err = rows.Scan(anys...)
		require.NoError(t, err)
		for i := 0; i < numCols; i++ {
		}
	}

	results := make([][]interface{}, numCols)
	for i := 0; i < numCols; i++ {
		results[i] = scanners[i].vals
	}
	return results
}

func TestUploadActivityJson(t *testing.T) {
	testFile := "testUploadActivityJson.duckdb"
	require := require.New(t)
	DBTest(t, testFile, func(db *sql.DB) {
		s1 := `{"Activity":{"id":10}}`
		s2 := `{"Activity":{"id":20}}`
		// interface for comparison with generically scanned data
		ss := []interface{}{s1, s2}
		require.NoError(UploadActivityJson(db, []StringJsonable{StringJsonable(s1), StringJsonable(s2)}))

		rows, err := db.Query(`select * from temp_etl;`)
		require.NoError(err)
		defer rows.Close()
		scanned := ScanAll(t, rows)
		require.EqualValues(scanned[0], ss)
		require.Equal(scanned[0], ss)

		rows, err = db.Query(`select * from etl;`)
		require.NoError(err)
		defer rows.Close()
		scanned = ScanAll(t, rows)
		// interface for comparison with generically scanned data
		expectedEtlIds := []interface{}{int64(1), int64(2)}
		expectedActivityIds := []interface{}{int64(10), int64(20)}
		require.EqualValues(scanned[0], expectedEtlIds)
		require.EqualValues(scanned[1], expectedActivityIds)
		require.EqualValues(scanned[2], ss)
	})
}

func TestMergeActivities(t *testing.T) {
	testFile := "testMergeActivities.duckdb"
	DBTest(t, testFile, func(db *sql.DB) {
		s1 := `{"Activity":{"id":10}}`
		s2 := `{"Activity":{"id":20}}`
		require.NoError(t, UploadActivityJson(db, []StringJsonable{StringJsonable(s1), StringJsonable(s2)}))

		//panicIfNotNil(MergeActivities(db))
	})
}

func TestScanColumns(t *testing.T) {
	require := require.New(t)
	DBTest(t, "", func(db *sql.DB) {
		row := db.QueryRow("create table t(c int64);")
		require.NoError(row.Err())
	})
}
