package common

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var dbConnection *sql.DB

func CreateDbConnection() *sql.DB {
	// connection
	cfg := mysql.Config{
		User:                 GetDotEnvVariable("DBUSER"),
		Passwd:               GetDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 GetDotEnvVariable("DBADDR"),
		DBName:               GetDotEnvVariable("DBNAME"),
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
