package intsql

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/stevenpelley/strava-snowflake/internal/util"
	"github.com/stretchr/testify/require"
)

type StringJsonable string

func (sj StringJsonable) ToJson() string {
	return string(sj)
}

func ToJsonables(slice []StringJsonable) []util.Jsonable {
	s := make([]util.Jsonable, 0, len(slice))
	for _, sj := range slice {
		s = append(s, sj)
	}
	return s
}

// scan the rows column-wise into the slices pointed to by ...any
// The arguments of Scan are constructed according to the types of the pointed slices.
// the inputs may also be a slice pointer converted to an interface{}
func ScanColumns(rows *sql.Rows, slicesOut ...any) error {
	numCols := len(slicesOut)
	slicePointerValues := make([]reflect.Value, numCols)
	scanPointerValues := make([]reflect.Value, numCols)
	scanArgs := make([]interface{}, numCols)

	for i, s := range slicesOut {
		t := reflect.TypeOf(s)
		// unwrap an interface{} into its element
		if t.Kind() == reflect.Interface {
			t = t.Elem()
		}

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

func ScanColumnsAndCompare(t *testing.T, rows *sql.Rows, expectedValsSlices ...any) {
	require := require.New(t)
	scanIfaces := make([]interface{}, 0)
	scanSlicePointers := make([]reflect.Value, 0)
	for _, s := range expectedValsSlices {
		t := reflect.TypeOf(s)
		require.Equal(reflect.Slice, t.Kind())
		// construct a slice of this elem type and then get the pointer to it to pass to scan
		slicePointerVal := reflect.New(t)
		scanSlicePointers = append(scanSlicePointers, slicePointerVal)

		ifaceVal := slicePointerVal.Interface()
		scanIfaces = append(scanIfaces, ifaceVal)
	}

	err := ScanColumns(rows, scanIfaces...)
	require.NoError(err)

	// compare the resulting scan
	//for idx, v := range expectedValsIfaces {
	for idx, v := range expectedValsSlices {
		i := reflect.Indirect(scanSlicePointers[idx]).Interface()
		require.EqualValues(v, i)
	}
}

func HelperTestUploadActivityJson(t *testing.T, sdb StravaDatabase) {
	require := require.New(t)
	s1 := `{"Activity":{"id":10}}`
	s2 := `{"Activity":{"id":20}}`
	// interface for comparison with generically scanned data
	ss := []string{s1, s2}
	require.NoError(sdb.UploadActivityJson(ToJsonables(
		[]StringJsonable{StringJsonable(s1), StringJsonable(s2)})))

	rows, err := sdb.DB().Query(`select * from temp_etl;`)
	require.NoError(err)
	defer rows.Close()
	ScanColumnsAndCompare(t, rows, ss)

	rows, err = sdb.DB().Query(`select * from etl;`)
	require.NoError(err)
	defer rows.Close()

	expectedEtlIds := []int64{int64(1), int64(2)}
	expectedActivityIds := []int64{int64(10), int64(20)}
	ScanColumnsAndCompare(t, rows, expectedEtlIds, expectedActivityIds, ss)
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
	sdb StravaDatabase,
	expectedMergeEffects expectedMergeEffects) {
	require := require.New(t)
	require.NoError(sdb.MergeActivities())

	// assert activities
	rows, err := sdb.DB().Query(`
		select etl_id, activity_id, activity.id, streamset.watts.resolution
		from activities
		order by activity_id;`)
	require.NoError(err)
	defer rows.Close()

	ScanColumnsAndCompare(
		t,
		rows,
		expectedMergeEffects.etlIds,
		expectedMergeEffects.activityIds,
		expectedMergeEffects.activityIdsFromDoc,
		expectedMergeEffects.resolutions)

	// assert streams
	rows, err = sdb.DB().Query(`
		select etl_id, activity_id, time, watts
		from streams
		order by activity_id, time;`)
	require.NoError(err)
	defer rows.Close()

	ScanColumnsAndCompare(
		t,
		rows,
		expectedMergeEffects.streamsEtlIds,
		expectedMergeEffects.streamsActivityIds,
		expectedMergeEffects.streamsTimes,
		expectedMergeEffects.streamsWatts)
}

func HelperTestMergeActivities(t *testing.T, sdb StravaDatabase) {
	require := require.New(t)
	s1 := `{"Activity":{"id":10},"StreamSet":{"watts":{"resolution":"asdf","data":[0, 100, 150]},"time":{"data":[0, 1, 2]}}}`
	s2 := `{"Activity":{"id":20},"StreamSet":{"watts":{"data":[200, 300, 150]},"time":{"data":[0, 1, 2]}}}`
	require.NoError(sdb.UploadActivityJson(ToJsonables(
		[]StringJsonable{StringJsonable(s1), StringJsonable(s2)})))

	mergeActivitiesAndAssertEffects(t, sdb, expectedMergeEffects{
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
	mergeActivitiesAndAssertEffects(t, sdb, expectedMergeEffects{
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
	require.NoError(sdb.UploadActivityJson(ToJsonables(
		[]StringJsonable{StringJsonable(s1), StringJsonable(s3)})))

	mergeActivitiesAndAssertEffects(t, sdb, expectedMergeEffects{
		etlIds:             []int64{3, 2, 4},
		activityIds:        []int64{10, 20, 30},
		activityIdsFromDoc: []int64{10, 20, 30},
		resolutions:        []sql.NullString{{String: "asdf", Valid: true}, {}, {}},

		streamsEtlIds:      []int64{3, 3, 3, 2, 2, 2, 4, 4, 4},
		streamsActivityIds: []int64{10, 10, 10, 20, 20, 20, 30, 30, 30},
		streamsTimes:       []int64{0, 1, 2, 0, 1, 2, 0, 1, 2},
		streamsWatts:       []int64{0, 100, 150, 200, 300, 150, 500, 100, 150},
	})
}
