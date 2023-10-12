package intsql

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

type StringJsonable string

func (sj StringJsonable) ToJson() string {
	return string(sj)
}

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

func DBTest(t *testing.T, testFile string, closeDb closeDb, f func(db *sql.DB)) {
	deleteTestFile(t, testFile)
	if closeDb == Close {
		defer deleteTestFile(t, testFile)
	}
	db, err := OpenDB(testFile)
	require.NoError(t, err)
	defer db.Close()

	require.NoError(t, InitAndValidateSchema(db))
	f(db)
}

// scan the rows column-wise into the slices pointed to by ...any
// The arguments of Scan are constructed according to the types of the pointed slices
func ScanColumns(rows *sql.Rows, slicesOut ...any) error {
	numCols := len(slicesOut)
	slicePointerValues := make([]reflect.Value, numCols)
	scanPointerValues := make([]reflect.Value, numCols)
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
		slicePointerValues[i] = reflect.ValueOf(s)

		elem2 := elem1.Elem()
		pointerValue := reflect.New(elem2)
		scanPointerValues[i] = pointerValue
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
			pointerValue := scanPointerValues[i]
			slicePointerValue := slicePointerValues[i]
			newSliceValue := reflect.Append(slicePointerValue.Elem(), pointerValue.Elem())
			reflect.Indirect(slicePointerValue).Set(newSliceValue)
		}
	}

	return nil
}

func TestUploadActivityJson(t *testing.T) {
	testFile := ""
	require := require.New(t)
	DBTest(t, testFile, Close, func(db *sql.DB) {
		s1 := `{"Activity":{"id":10}}`
		s2 := `{"Activity":{"id":20}}`
		// interface for comparison with generically scanned data
		ss := []string{s1, s2}
		require.NoError(UploadActivityJson(db, []StringJsonable{StringJsonable(s1), StringJsonable(s2)}))

		rows, err := db.Query(`select * from temp_etl;`)
		require.NoError(err)
		defer rows.Close()
		var scanned []string
		err = ScanColumns(rows, &scanned)
		require.NoError(err)
		require.EqualValues(scanned, ss)

		rows, err = db.Query(`select * from etl;`)
		require.NoError(err)
		defer rows.Close()
		var scanned0 []int64
		var scanned1 []int64
		var scanned2 []string
		err = ScanColumns(rows, &scanned0, &scanned1, &scanned2)
		require.NoError(err)

		expectedEtlIds := []int64{int64(1), int64(2)}
		expectedActivityIds := []int64{int64(10), int64(20)}
		require.EqualValues(scanned0, expectedEtlIds)
		require.EqualValues(scanned1, expectedActivityIds)
		require.EqualValues(scanned2, ss)
	})
}

type expectedMergeEffects struct {
	etlIds             []int64
	activityIds        []int64
	activityIdsFromDoc []int64
	resolutions        []sql.NullString

	streamsEtlIds      []int64
	streamsActivityIds []int64
	streamsTimes       []int64
	streamsWatts       []int64
}

func mergeActivitiesAndAssertEffects(
	t *testing.T,
	db *sql.DB,
	expectedMergeEffects expectedMergeEffects) {
	require := require.New(t)
	require.NoError(MergeActivities(db))

	// assert activities
	rows, err := db.Query(`
		select etl_id, activity_id, activity.id, streamset.watts.resolution
		from activities
		order by activity_id;`)
	require.NoError(err)
	defer rows.Close()

	var etlIds []int64
	var activityIds []int64
	var activityIdsFromDoc []int64
	var resolutions []sql.NullString
	err = ScanColumns(rows, &etlIds, &activityIds, &activityIdsFromDoc, &resolutions)
	require.NoError(err)

	require.EqualValues(expectedMergeEffects.etlIds, etlIds)
	require.EqualValues(expectedMergeEffects.activityIds, activityIds)
	require.EqualValues(expectedMergeEffects.activityIdsFromDoc, activityIdsFromDoc)
	require.EqualValues(expectedMergeEffects.resolutions, resolutions)

	// assert streams
	rows, err = db.Query(`
		select etl_id, activity_id, time, watts
		from streams
		order by activity_id, time;`)
	require.NoError(err)
	defer rows.Close()

	etlIds = nil
	activityIds = nil
	var times []int64
	var watts []int64
	err = ScanColumns(rows, &etlIds, &activityIds, &times, &watts)
	require.NoError(err)

	require.EqualValues(expectedMergeEffects.streamsEtlIds, etlIds)
	require.EqualValues(expectedMergeEffects.streamsActivityIds, activityIds)
	require.EqualValues(expectedMergeEffects.streamsTimes, times)
	require.EqualValues(expectedMergeEffects.streamsWatts, watts)
}

func TestMergeActivities(t *testing.T) {
	testFile := ""
	require := require.New(t)

	DBTest(t, testFile, Close, func(db *sql.DB) {
		s1 := `{"Activity":{"id":10},"StreamSet":{"watts":{"resolution":"asdf","data":[0, 100, 150]},"time":{"data":[0, 1, 2]}}}`
		s2 := `{"Activity":{"id":20},"StreamSet":{"watts":{"data":[200, 300, 150]},"time":{"data":[0, 1, 2]}}}`
		require.NoError(UploadActivityJson(db, []StringJsonable{StringJsonable(s1), StringJsonable(s2)}))

		mergeActivitiesAndAssertEffects(t, db, expectedMergeEffects{
			etlIds:             []int64{1, 2},
			activityIds:        []int64{10, 20},
			activityIdsFromDoc: []int64{10, 20},
			resolutions:        []sql.NullString{{String: "asdf", Valid: true}, {}},

			streamsEtlIds:      []int64{1, 1, 1, 2, 2, 2},
			streamsActivityIds: []int64{10, 10, 10, 20, 20, 20},
			streamsTimes:       []int64{0, 1, 2, 0, 1, 2},
			streamsWatts:       []int64{0, 100, 150, 200, 300, 150},
		})

		// should be idempotent
		mergeActivitiesAndAssertEffects(t, db, expectedMergeEffects{
			etlIds:             []int64{1, 2},
			activityIds:        []int64{10, 20},
			activityIdsFromDoc: []int64{10, 20},
			resolutions:        []sql.NullString{{String: "asdf", Valid: true}, {}},

			streamsEtlIds:      []int64{1, 1, 1, 2, 2, 2},
			streamsActivityIds: []int64{10, 10, 10, 20, 20, 20},
			streamsTimes:       []int64{0, 1, 2, 0, 1, 2},
			streamsWatts:       []int64{0, 100, 150, 200, 300, 150},
		})

		// now insert additional activities, some new and some duplicates
		s3 := `{"Activity":{"id":30},"StreamSet":{"watts":{"data":[500, 100, 150]},"time":{"data":[0, 1, 2]}}}`
		require.NoError(UploadActivityJson(db, []StringJsonable{StringJsonable(s1), StringJsonable(s3)}))

		mergeActivitiesAndAssertEffects(t, db, expectedMergeEffects{
			etlIds:             []int64{3, 2, 4},
			activityIds:        []int64{10, 20, 30},
			activityIdsFromDoc: []int64{10, 20, 30},
			resolutions:        []sql.NullString{{String: "asdf", Valid: true}, {}, {}},

			streamsEtlIds:      []int64{3, 3, 3, 2, 2, 2, 4, 4, 4},
			streamsActivityIds: []int64{10, 10, 10, 20, 20, 20, 30, 30, 30},
			streamsTimes:       []int64{0, 1, 2, 0, 1, 2, 0, 1, 2},
			streamsWatts:       []int64{0, 100, 150, 200, 300, 150, 500, 100, 150},
		})
	})
}

func TestScanColumns(t *testing.T) {
	require := require.New(t)
	DBTest(t, "", Close, func(db *sql.DB) {
		row := db.QueryRow("create table t(c1 int64, c2 string);")
		require.NoError(row.Err())

		row = db.QueryRow("insert into t values (1, 'asdf'), (2, 'qwer');")
		require.NoError(row.Err())

		rows, err := db.Query("select * from t;")
		require.NoError(err)
		defer rows.Close()

		var ints []int64
		var strings []string
		err = ScanColumns(rows, &ints, &strings)
		require.NoError(err)

		require.EqualValues(ints, []int64{1, 2})
		require.EqualValues(strings, []string{"asdf", "qwer"})
	})
}
