package main

// https://github.com/microsoft/go-mssqldb

// go run main.go 15789

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

var (
	debug    	= flag.Bool("debug", false, "enable debugging")
	databases	= flag.Bool("databases", false, "retrieve database names from sql instance")
	password 	= flag.String("password", "", "the database password")
	port     	= flag.Int("port", 1433, "the database port")
	server   	= flag.String("server", "localhost", "the database server")
	user     	= flag.String("user", "sa", "the database user")
	query    	= flag.String("query", "", "custom SQL query to execute")
	timeout  	= flag.Duration("timeout", 30*time.Second, "connection timeout")
	version  	= flag.Bool("version", false, "show version information")
)

func main() {

	flag.Parse()

	if *version {
		fmt.Println("prusk-sql v1.0.0")
		fmt.Println("A command-line utility for connecting to and querying SQL Server")
		return
	}

	// Get password from environment if not provided
	actualPassword := *password
	if actualPassword == "" {
		actualPassword = os.Getenv("SQL_PASSWORD")
		if actualPassword == "" {
			log.Fatal("Password must be provided via -password flag or SQL_PASSWORD environment variable")
		}
	}

	if *debug {
		fmt.Printf(" password:*****\n")
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" timeout:%s\n", *timeout)
	}

	fmt.Printf("OK, trying to connect to SQL...\n")
	fmt.Printf("\n")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// this doesn't seem to be throwing an error if it cannot connect...
	// because sql.Open() does not open a connection to a database
	// it just validates the arguments
	// Ping() is used to test the connection
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, actualPassword, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	// Test connection with context
	if err = conn.PingContext(ctx); err != nil {
		log.Fatal("Ping connection failed:", err.Error())
	}

	// get sql name
	fmt.Printf("Connected to SQL! Let's get the instance name...\n")

	row1 := conn.QueryRowContext(ctx, "SELECT @@SERVERNAME")
	var sqlname string
	err = row1.Scan(&sqlname)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	fmt.Printf(sqlname)
	fmt.Printf("\n")
	fmt.Printf("\n")

	// get sql version
	fmt.Printf("Let's get the version running...\n")

	row2 := conn.QueryRowContext(ctx, "SELECT @@VERSION")
	var sqlversion string
	err = row2.Scan(&sqlversion)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	fmt.Printf(sqlversion)
	fmt.Printf("\n")
	fmt.Printf("\n")


	if *databases {
		// get the databases in the sql instance
		fmt.Printf("What databases are in the instance?\n")

		rows, err := conn.QueryContext(ctx, "SELECT [name] FROM sys.databases")
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var databasename string

			err := rows.Scan(&databasename)
			if err != nil {
				log.Fatal("Scan failed:", err.Error())
			}

			fmt.Printf(databasename)
			fmt.Printf("\n")
		}
		if err = rows.Err(); err != nil {
			log.Fatal("Rows error:", err.Error())
		}
		fmt.Printf("\n")
	}

	if *query != "" {
		fmt.Printf("Executing custom query: %s\n", *query)
		fmt.Printf("\n")

		rows, err := conn.QueryContext(ctx, *query)
		if err != nil {
			log.Fatal("Query failed:", err.Error())
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			log.Fatal("Columns failed:", err.Error())
		}

		if len(columns) > 0 {
			// Print column headers
			for i, col := range columns {
				if i > 0 {
					fmt.Printf("\t")
				}
				fmt.Printf(col)
			}
			fmt.Printf("\n")

			// Create slice to hold row values
			values := make([]interface{}, len(columns))
			scanArgs := make([]interface{}, len(values))
			for i := range values {
				scanArgs[i] = &values[i]
			}

			// Print rows
			for rows.Next() {
				err = rows.Scan(scanArgs...)
				if err != nil {
					log.Fatal("Scan failed:", err.Error())
				}

				for i, val := range values {
					if i > 0 {
						fmt.Printf("\t")
					}
					if val != nil {
						fmt.Printf("%v", val)
					} else {
						fmt.Printf("NULL")
					}
				}
				fmt.Printf("\n")
			}
		} else {
			fmt.Printf("Query executed successfully.\n")
		}

		if err = rows.Err(); err != nil {
			log.Fatal("Rows error:", err.Error())
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	fmt.Printf("See ya!\n")

}
