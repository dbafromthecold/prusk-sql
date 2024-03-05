package main

import (
	"flag"
	"fmt"
)

var (
	debug     = flag.Bool("debug", false, "enable debugging")
	databases = flag.Bool("databases", false, "retrieve database names from SQL instance")
	password  = flag.String("password", "Testing1122", "the database password")
	port      = flag.Int("port", 1433, "the database port")
	server    = flag.String("server", "localhost", "the database server")
	user      = flag.String("user", "sa", "the database user")
)

func initializeConfig() {
	if *debug {
		fmt.Printf("password:%s\n", *password)
		fmt.Printf("port:%d\n", *port)
		fmt.Printf("server:%s\n", *server)
		fmt.Printf("user:%s\n", *user)
	}
}