package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var dbConnection *sql.DB

func createDbConnection() *sql.DB {
	// connection
	cfg := mysql.Config{
		User:                 getDotEnvVariable("DBUSER"),
		Passwd:               getDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 getDotEnvVariable("DBADDR"),
		DBName:               getDotEnvVariable("DBNAME"),
		AllowNativePasswords: true,
	}

	fmt.Printf("%+v\n", cfg.FormatDSN())
	// Get a database handle.
	var err error
	dbConnection, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := dbConnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return dbConnection
}
