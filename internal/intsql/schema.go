// functions to create and verify our schema.  Should rely only on database/sql
package intsql

import (
	"database/sql"
	"fmt"
	"reflect"
)

const etlSequence = "etl_seq"
const etlTable = "etl"
const tempEtlTable = "temp_etl"
const activitiesTable = "activities"
const streamsTable = "streams"
const activityIdsToUpdateView = "activity_ids_to_update"
const newActivitiesView = "new_activities"
const newStreamsView = "new_streams"

func InitAndValidateSchema(db *sql.DB) error {
	err := InitSchema(db)
	if err != nil {
		return fmt.Errorf("error initializing data schema: %w", err)
	}

	err = ValidateSchema(db)
	if err != nil {
		return fmt.Errorf("error validating schema: %w", err)
	}

	return nil
}

// verify that the expected tables exist in the database.
// Databases may implicitly create a new db instead of providing an error
// so this should be called when we expect an existing db.
func ValidateSchema(db *sql.DB) error {
	sqlText := `select table_name from information_schema.tables where table_catalog=current_database() and table_schema=current_schema() and table_name in ('streams', 'activities', 'etl', 'temp_etl') order by table_name;`
	rows, err := db.Query(sqlText)
	if err != nil {
		return err
	}
	defer rows.Close()
	results := make([]string, 0, 4)
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return err
		}
		results = append(results, s)
	}
	rows.Close()

	expected := []string{"activities", "etl", "streams", "temp_etl"}
	if !reflect.DeepEqual(expected, results) {
		return fmt.Errorf(
			"unexpected existing tables.  Expected 'activities', 'etl', 'streams', 'temp_etl'.  Found: %v",
			results)
	}

	sqlText = `select function_name from duckdb_functions()
	where not internal and
	database_name = current_database()
	order by function_name;`
	rows, err = db.Query(sqlText)
	if err != nil {
		return err
	}
	defer rows.Close()
	results = make([]string, 0, 3)
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return err
		}
		results = append(results, s)
	}
	rows.Close()

	expected = []string{"expand_streams"}
	if !reflect.DeepEqual(expected, results) {
		return fmt.Errorf(
			"unexpected existing macros.  Expected %v.  Found: %v",
			expected, results)
	}

	return nil
}

func InitSchema(db *sql.DB) error {
	err := CreateMacros(db)
	if err != nil {
		return fmt.Errorf("error creating macros: %w", err)
	}
	err = CreateETLSequence(db)
	if err != nil {
		return fmt.Errorf("error creating etl sequence: %w", err)
	}
	err = CreateTempETLTable(db)
	if err != nil {
		return fmt.Errorf("error creating temp_etl table: %w", err)
	}
	err = CreateETLTable(db)
	if err != nil {
		return fmt.Errorf("error creating etl table: %w", err)
	}
	err = CreateActivitiesTable(db)
	if err != nil {
		return fmt.Errorf("error creating activities table: %w", err)
	}
	err = CreateStreamsTable(db)
	if err != nil {
		return fmt.Errorf("error creating streams table: %w", err)
	}
	err = CreateActivityIdsToUpdateView(db)
	if err != nil {
		return fmt.Errorf("create %v: %w", activityIdsToUpdateView, err)
	}
	err = CreateDedupedEtlView(db)
	if err != nil {
		return fmt.Errorf("create dedupted_etl: %w", err)
	}
	err = CreateNewActivitiesView(db)
	if err != nil {
		return fmt.Errorf("create new_activities: %w", err)
	}
	err = CreateNewStreamsView(db)
	if err != nil {
		return fmt.Errorf("create new_streams: %w", err)
	}
	return nil
}

func CreateETLSequence(db *sql.DB) error {
	row := db.QueryRow(fmt.Sprintf("create sequence if not exists \"%v\";", etlSequence))
	return row.Err()
}

func CreateTempETLTable(db *sql.DB) error {
	row := db.QueryRow(
		fmt.Sprintf("create table if not exists \"%v\" (json json);",
			tempEtlTable))
	return row.Err()
}

func CreateETLTable(db *sql.DB) error {
	row := db.QueryRow(
		fmt.Sprintf("create table if not exists \"%v\" (etl_id int64 primary key default nextval('%v'), activity_id int64, json json);",
			etlTable, etlSequence))
	return row.Err()
}

// TODO: add primary key back once duckdb supports transaction updates to primary key (see indexes doc)
func CreateActivitiesTable(db *sql.DB) error {
	sqlText := fmt.Sprintf(`
create or replace table %v
(
activity_id int64
--	primary key
	,
etl_id int64,
activity ACTIVITY_T,
streamset STREAMSET_NODATA_T,
);
`, activitiesTable)
	row := db.QueryRow(sqlText)
	return row.Err()
}

// TODO: add primary key back once duckdb supports transaction updates to primary key (see indexes doc)
func CreateStreamsTable(db *sql.DB) error {
	sqlText := fmt.Sprintf(`
create or replace table %v
(
activity_id int64,
etl_id int64,
time int64,
watts int64,
heartrate int64,
velocity_smooth double,
grade_smooth double,
cadence int64,
distance double,
altitude double,
latlong double[],
temp int64,
moving boolean,
--primary key (activity_id, time) -- compound primary key
);
`, streamsTable)
	row := db.QueryRow(sqlText)
	return row.Err()
}

func CreateActivityIdsToUpdateView(db *sql.DB) error {
	sqlText := fmt.Sprintf(`
create or replace view %v as
select activity_id
from activities
where exists (
        select 1
        from etl
        where activities.activity_id = etl.activity_id and activities.etl_id < etl.etl_id
        );`,
		activityIdsToUpdateView)
	row := db.QueryRow(sqlText)
	return row.Err()
}

func CreateDedupedEtlView(db *sql.DB) error {
	sqlText := `
create or replace view deduped_etl as
select *
from etl
where etl_id = (select max(etl_id) from etl as etl2 where etl.activity_id = etl2.activity_id);`
	row := db.QueryRow(sqlText)
	return row.Err()
}

func CreateNewActivitiesView(db *sql.DB) error {
	sqlText := fmt.Sprintf(`
create or replace view %v as
select 
	activity_id,
	etl_id,
	try_cast(json_extract("json", '$.Activity') as ACTIVITY_T),
	try_cast(json_extract("json", '$.StreamSet') as STREAMSET_NODATA_T)
from deduped_etl
where coalesce(etl_id > (select max(etl_id) from activities where deduped_etl.activity_id = activities.activity_id), true)
	;`,
		newActivitiesView)
	row := db.QueryRow(sqlText)
	return row.Err()
}

func CreateNewStreamsView(db *sql.DB) error {
	sqlText := fmt.Sprintf(`
create or replace view %v as
select 
	activity_id,
	etl_id,
	s.*
from deduped_etl, expand_streams(try_cast(json_extract(deduped_etl."json", '$.StreamSet') as STREAMSET_T)) as s
where coalesce(etl_id > (select max(etl_id) from streams where deduped_etl.activity_id = streams.activity_id), true)
	;`,
		newStreamsView)
	row := db.QueryRow(sqlText)
	return row.Err()
}

func createOrReplaceType(db *sql.DB, typeName string, sqlText string) error {
	row := db.QueryRow(fmt.Sprintf(`drop type if exists %v;`, typeName))
	if row.Err() != nil {
		return fmt.Errorf("error dropping type %v: %w", typeName, row.Err())
	}
	row = db.QueryRow(sqlText)
	if row.Err() != nil {
		return fmt.Errorf("error creating type %v: %w", typeName, row.Err())
	}

	return nil
}

func CreateMacros(db *sql.DB) error {
	createOrReplaceType(db, "ACTIVITY_T", `create type ACTIVITY_T as struct(achievement_count bigint, athlete struct(id bigint), athlete_count bigint, average_speed double, average_watts double, comment_count bigint, commute boolean, device_watts boolean, distance double, elapsed_time bigint, elev_high double, elev_low double, end_latlng double[], external_id varchar, flagged boolean, gear_id varchar, has_kudoed boolean, id bigint, kilojoules double, kudos_count bigint, manual boolean, "map" struct(id varchar, summary_polyline varchar), max_speed double, max_watts bigint, moving_time bigint, "name" varchar, photo_count bigint, private boolean, sport_type varchar, start_date timestamp, start_date_local timestamp, start_latlng double[], timezone varchar, total_elevation_gain double, total_photo_count bigint, trainer boolean, "type" varchar, upload_id bigint, upload_id_str bigint, weighted_average_watts bigint, workout_type bigint);`)
	createOrReplaceType(db, "STREASET_T", `create type STREAMSET_T as struct(altitude struct("data" double[], original_size bigint, resolution varchar, series_type varchar), cadence struct("data" bigint[], original_size bigint, resolution varchar, series_type varchar), distance struct("data" double[], original_size bigint, resolution varchar, series_type varchar), grade_smooth struct("data" double[], original_size bigint, resolution varchar, series_type varchar), heartrate struct("data" bigint[], original_size bigint, resolution varchar, series_type varchar), latlng struct("data" double[][], original_size bigint, resolution varchar, series_type varchar), moving struct("data" boolean[], original_size bigint, resolution varchar, series_type varchar), "time" struct("data" bigint[], original_size bigint, resolution varchar, series_type varchar), velocity_smooth struct("data" double[], original_size bigint, resolution varchar, series_type varchar), watts struct("data" bigint[], original_size bigint, resolution varchar, series_type varchar), "temp" struct("data" bigint[], original_size bigint, resolution varchar, series_type varchar));`)
	createOrReplaceType(db, "STREAMSET_NODATA_T", `create type STREAMSET_NODATA_T as struct(altitude struct(original_size bigint, resolution varchar, series_type varchar), cadence struct(original_size bigint, resolution varchar, series_type varchar), distance struct(original_size bigint, resolution varchar, series_type varchar), grade_smooth struct(original_size bigint, resolution varchar, series_type varchar), heartrate struct(original_size bigint, resolution varchar, series_type varchar), latlng struct(original_size bigint, resolution varchar, series_type varchar), moving struct(original_size bigint, resolution varchar, series_type varchar), "time" struct(original_size bigint, resolution varchar, series_type varchar), velocity_smooth struct(original_size bigint, resolution varchar, series_type varchar), watts struct(original_size bigint, resolution varchar, series_type varchar), "temp" struct(original_size bigint, resolution varchar, series_type varchar));`)

	sqlText := `create or replace macro expand_streams(c) as table
		select
		  unnest(c['time']['data']) as time,
		  unnest(c['watts']['data']) as watts,
		  unnest(c['heartrate']['data']) as heartrate,
		  unnest(c['velocity_smooth']['data']) as velocity_smooth,
		  unnest(c['grade_smooth']['data']) as grade_smooth,
		  unnest(c['cadence']['data']) as cadence,
		  unnest(c['distance']['data']) as distance,
		  unnest(c['altitude']['data']) as altitude,
		  unnest(c['latlng']['data']) as latlng,
		  unnest(c['temp']['data']) as temp,
		  unnest(c['moving']['data']) as moving,
		;`
	row := db.QueryRow(sqlText)
	if row.Err() != nil {
		return fmt.Errorf("error creating macro expand_streams: %w", row.Err())
	}

	return nil
}
