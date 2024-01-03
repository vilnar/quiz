package appdb

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func GetCountRowsInTable(db *sql.DB, tableName string) int {
	var count int

	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of rows are %s\n", count)
	return count
}

func Placeholders(n int) string {
	ps := make([]string, n)
	for i := 0; i < n; i++ {
		ps[i] = "?"
	}
	return strings.Join(ps, ",")
}

func IdsToArgs(ids []int64) []interface{} {
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return args
}
