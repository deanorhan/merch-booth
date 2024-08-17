package main

import (
	"database/sql"
	"log"
)

var conn *sql.DB

func init() {
	var err error

	conn, err = sql.Open("sqlite3", "./merch.db")
	if err != nil {
		log.Panic(err)
	}
}

func GetDBConnection() *sql.DB {
	return conn
}
