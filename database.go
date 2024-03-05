package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/microsoft/go-mssqldb"
)

func openDatabaseConnection() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf("connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getSQLInstanceName(conn *sql.DB) {
	// similar to original code, just ensure it's wrapped in a function
}

func getSQLVersion(conn *sql.DB) {
	// similar to original code, just ensure it's wrapped in a function
}

func getDatabases(conn *sql.DB) {
	// similar to original code, just ensure it's wrapped in a function
}