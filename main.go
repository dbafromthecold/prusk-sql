package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()
	initializeConfig()
	fmt.Printf("OK, trying to connect to SQL...\n\n")

	conn, err := openDatabaseConnection()
	if err != nil {
		log.Fatalf("Open connection failed: %v", err)
	}
	defer conn.Close()

	getSQLInstanceName(conn)
	getSQLVersion(conn)
	if *databases {
		getDatabases(conn)
	}
	fmt.Printf("\nSee ya!\n")
}
