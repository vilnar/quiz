package appdb

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"quiz/internal/common"
)

var dbConnection *sql.DB

func CreateDbConnection() *sql.DB {
	// connection
	cfg := mysql.Config{
		User:                 common.GetDotEnvVariable("DBUSER"),
		Passwd:               common.GetDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 common.GetDotEnvVariable("DBADDR"),
		DBName:               common.GetDotEnvVariable("DBNAME"),
		AllowNativePasswords: true,
	}

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
	log.Println("Connected!")
	return dbConnection
}
