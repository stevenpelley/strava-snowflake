package main

import (
	"fmt"
	"log"

	"github.com/stevenpelley/strava-snowflake/internal/intsql"
	"github.com/stevenpelley/strava-snowflake/internal/strava"
)

func main() {
	strava.InitLogging("test.log")

	db, err := intsql.OpenDB("")
	if err != nil {
		log.Panicf("error opening db file: %v", err)
	}
	defer db.Close()

	row := db.QueryRow("select version();")
	if row.Err() != nil {
		log.Panicf("error reading version: %v", row.Err())
	}
	var version string
	err = row.Scan(&version)
	if err != nil {
		log.Panicf("error scanning version: %v", err)
	}
	fmt.Printf("version: %v\n", version)

	//_, err = db.Query("create or replace temp table temptable1 (n int64);")
	if err != nil {
		log.Panicf("error creating table from db: %v", err)
	}

	_, err = db.Query("create or replace temp table temptable1 (json json);")
	//_, err = db.Query("create or replace temp table temptable1 (n int64);")
	if err != nil {
		log.Panicf("error creating table from db: %v", err)
	}

	////conn, err := db.Conn(context.Background())
	////if err != nil {
	////	log.Panicf("error getting a connection: %v", err)
	////}
	////defer conn.Close()

	////_, err = conn.QueryContext(context.Background(), "create or replace temp table temptable2 (json json);")
	////if err != nil {
	////	log.Panicf("error creating table from conn: %v", err)
	////}
}
