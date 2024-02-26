package appdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"quiz/internal/common"
)

var dbConnection *sql.DB

func CreateDbConnection() *sql.DB {
	var err error
	dbConnection, err = sql.Open("sqlite3", common.GetDbPath())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := dbConnection.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	return dbConnection
}
