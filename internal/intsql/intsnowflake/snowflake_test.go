package intsnowflake

import (
	"context"
	"fmt"
	"testing"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stretchr/testify/require"
)

const testSequence = `"STRAVA"."ACTIVITIES_TEST"."ETL_SEQUENCE_TEST"`
const testTable = `"STRAVA"."ACTIVITIES_TEST"."ETL_TEST"`

func DBTest(t *testing.T, f func(sdb *SFStrava)) {
	require := require.New(t)
	sdb := New("../../../snowflake_config.json", testTable)
	err := sdb.OpenDB()
	require.NoError(err)
	defer sdb.Close()

	_, err = sdb.DB().ExecContext(
		context.Background(),
		fmt.Sprintf(`create or replace sequence %v;`, testSequence))
	require.NoError(err)

	_, err = sdb.DB().ExecContext(
		context.Background(),
		fmt.Sprintf(
			`create or replace transient table %v (etl_id number default
			%v.nextval, data variant);`,
			testTable,
			testSequence))
	require.NoError(err)

	defer sdb.DB().ExecContext(context.Background(), fmt.Sprintf(`drop %v;`, testTable))

	require.NoError(sdb.InitAndValidateSchema())
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
		require := require.New(t)
		s1 := `{"Activity":{"id":10}}`
		s2 := `{"Activity":{"id":20}}`
		// interface for comparison with generically scanned data
		ss := []string{s1, s2}
		require.NoError(sdb.UploadActivityJson(intsql.ToJsonables(
			[]intsql.StringJsonable{intsql.StringJsonable(s1), intsql.StringJsonable(s2)})))

		rows, err := sdb.DB().Query(
			fmt.Sprintf(
				`select data::string from %v;`,
				testTable))
		require.NoError(err)
		defer rows.Close()
		intsql.ScanColumnsAndCompare(t, rows, ss)
	})
}

func TestMergeActivities(t *testing.T) {
	DBTest(t, func(sdb *SFStrava) {
		intsql.HelperTestMergeActivities(
			t,
			sdb,

			fmt.Sprintf(`
			select
			  etl_id,
			  data:"Activity":id::number as activity_id,
			  data:"Activity":id::number as activity_id2,
			  data:"StreamSet":"watts":"resolution"::string as resolution
			from %[1]v as outer
			where etl_id = (
				select max(etl_id)
				from %[1]v as inner
				where inner.data:"Activity":id::number = outer.data:"Activity":id::number)
			order by activity_id
			;`, testTable),

			fmt.Sprintf(`
			with my_etl as (
			select
				*
			from %[1]v as outer
			where etl_id = (
				select max(etl_id)
				from %[1]v as inner
				where inner.data:"Activity":id::number = outer.data:"Activity":id::number)
			  ), times as (
				select
				etl_id,
				data:"Activity":"id"::number as activity_id,
				f.value as time,
				f.index as idx
				from my_etl, table(flatten(data:"StreamSet":"time":"data")) as f
			  ), watts as (
				select
				etl_id,
				data:"Activity":"id"::number as activity_id,
				f.value as watts,
				f.index as idx
				from my_etl, table(flatten(data:"StreamSet":"watts":"data")) as f
			  )
			  select
				coalesce(times.etl_id, watts.etl_id) as etl_id,
				coalesce(times.activity_id, watts.activity_id) as activity_id,
				times.time,
				watts.watts
			  from times full outer join watts on times.etl_id = watts.etl_id and times.activity_id = watts.activity_id and times.idx = watts.idx
			  order by activity_id, time
			  ;`, testTable))
	})
}
