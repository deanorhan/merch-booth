package main

import (
	"database/sql"
	"log"
)

var conn *sql.DB

func Open() error {
	log.Print("Opening the DB")
	var err error
	conn, err = sql.Open("sqlite3", "./merch.db")
	return err
}

func Close() {
	log.Print("Closing the DB")
	conn.Close()
}

func GetDBConnection() *sql.DB {
	return conn
}
