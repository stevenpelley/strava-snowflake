# strava-snowflake
track and export personal Strava data.  Ingest and visualize in Snowflake.

copy github.com/stevenpelley/strava3golang/oauth_client_config.json and the resulting token.json here.
Also needs snowflake key pair credentials placed in snowflake_config (see template)

## Current State
passing tests for both duckdb and snowflake
downloads activities within time range from strava
skips retrieving streams for activities already in database or listed in command line options
stores and transforms data in duckdb.
python snowpark udfs organized, flatten_streams written and tested.  Additionally fills time gaps.
vectorized ewma done.  Surprisingly slow.
non-vectorized ewma done.  Much faster.  Also simpler.  And correlates rows through the lateral join.
Looked into managing stages and UDF artifacts.  Temporary (anonymous) UDFs created by snowpark do clean up the associated files, but only once the session is cleaned.  If the session is not explicitly closed this takes 24 hours.  Close your sessions, people (me).  Named UDFs do not automatically have their files cleaned up.  This is problematic because I "create or replace" via snowpark, leaking the stage files.  Conclusion: always know where UDF stage files go (e.g., put all of them in a single stage and don't put other files there) and then outer join files in the stage (LIST) against UDF artifact locations ("import" field of information_schema.functions).  Worksheet snippets below.
Created a simple pythong script in src/main.py to register the UDFs as permanent UDFs.

## TODO
* try out dbt for ETL.  This might drastically simplify materializing streams.
* dashboarding
  * simple dashboards should provide list of activities with summary stats, the ability to drill into each activity and provide more data, as well as a time-window-based training progress chart.
  * try out snowflake builtin dashboarding
  * try out some other dashboard options (superset, mesosphere)

## Random information

### Cleaning up UDF stage files worksheet:
```
list @udf_stage;
-- use the query id above to query those files in a select statement for joining
select * from table(result_scan('01b031a7-0604-d114-0001-305300036222'));


with funcs as (
    select
        function_catalog,
        function_schema,
        function_name,
        function_owner,
        argument_signature,
        data_type,
        function_language,
        created,
        last_altered,
        imp.value::string as import_value,
        substr(import_value, position('/', import_value)) as file_name,
        handler,
        packages
    from information_schema.functions, table(flatten(parse_json(imports))) as imp
    where 1=1
        and function_catalog = 'STRAVA'
        and function_language = 'PYTHON'
        and function_name not ilike 'snowpark_temp%'
), stage_files as (
  select
      *,
      substr("name", position('/', "name")) as file_name
  from table(result_scan('01b031a7-0604-d114-0001-305300036222'))
), joined as (
    select
        function_catalog,
        function_schema,
        function_name,
        import_value,
        funcs.file_name as funcs_file_name,
        stage_files."name" as stage_files_name,
        stage_files.file_name as stage_files_file_name
    from funcs full outer join stage_files on funcs.file_name = stage_files.file_name
)
select *
from joined
where function_name is null
;


remove @udf_stage/bf8643c74b62f4ba5d095beb4108a2a7c2f62191dc92106e13119af3594010f7/src.zip;
```