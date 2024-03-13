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

func getSQLInstanceName(conn *sql.DB) (string, error) {
	var instanceName string
	err := conn.QueryRow("SELECT @@SERVERNAME").Scan(&instanceName)
	if err != nil {
		return "", err
	}
	return instanceName, nil
}

func getDatabases(conn *sql.DB) ([]string, error) {
	rows, err := conn.Query("SELECT name FROM sys.databases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var databases []string
	for rows.Next() {
		var databaseName string
		err := rows.Scan(&databaseName)
		if err != nil {
			return nil, err
		}
		databases = append(databases, databaseName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return databases, nil
}

func runCustomQuery(conn *sql.DB, query string) ([]map[string]interface{}, error) {
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, err
		}

		rowData := make(map[string]interface{})
		for i, col := range columns {
			rowData[col] = values[i]
		}
		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func getSQLVersion(conn *sql.DB) (string, error) {
	var version string
	err := conn.QueryRow("SELECT @@VERSION").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

