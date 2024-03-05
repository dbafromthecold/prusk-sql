package main

// https://github.com/microsoft/go-mssqldb

// go run main.go 15789

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	_ "github.com/microsoft/go-mssqldb"
)

var (
	debug    	= flag.Bool("debug", false, "enable debugging")
	databases	= flag.Bool("databases", false, "retreive database names from sql instance")
	password 	= flag.String("password", "Testing1122", "the database password")
	port     	= flag.Int("port", 1433, "the database port")
	server   	= flag.String("server", "localhost", "the database server")
	user     	= flag.String("user", "sa", "the database user")
)

func main() {

	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}

	fmt.Printf("OK, trying to connect to SQL...\n")
	fmt.Printf("\n")

	// this doesn't seem to be throwing an error if it cannot connect...
	// because sql.Open() does not open a connection to a database
	// it just validates the arguments
	// Ping() is used to test the connection
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	} else {
		err = conn.Ping()
		if err != nil {
			log.Fatal("Open connection failed:", err.Error())
		}
	}
	defer conn.Close()

	// get sql name
	fmt.Printf("Connected to SQL! Let's get the instance name...\n")

	stmt1, err := conn.Prepare("select @@servername")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt1.Close()

	row1 := stmt1.QueryRow()
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

	stmt2, err := conn.Prepare("select @@VERSION")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt2.Close()

	row2 := stmt2.QueryRow()
	var sqlversion string
	err = row2.Scan(&sqlversion)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	fmt.Printf(sqlversion)
	fmt.Printf("\n")
	fmt.Printf("\n")


	if *databases{
		// get the databases in the sql instance
		fmt.Printf("What databases are in the instance?\n")

		stmt3, err := conn.Prepare("select [name] FROM sys.databases;")
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt3.Close()

		rows, err := stmt3.Query()

		for rows.Next() {
			var (databasename string)

			err := rows.Scan(&databasename)
			if err != nil {
				log.Fatal("Scan failed:", err.Error())
			}

			fmt.Printf(databasename)
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
	fmt.Printf("See ya!\n")

}
