package appdb

import (
	"database/sql"
	"fmt"
	"log"
)

func GetCountRowsInTable(db *sql.DB, tableName string) int {
	var count int

	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number of rows are %s\n", count)
	return count
}
