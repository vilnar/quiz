package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	// mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"quiz/internal/common"
)

func main() {
	fmt.Printf("run setup\n\n")

	isDrop := flag.Bool("drop", false, "drop db")
	flag.Parse()

	runMigrate(*isDrop)
}

func runMigrate(isDrop bool) {
	// log.Printf("DSN %+v\n", cfg.FormatDSN())
	//Create a new SQLite database
	db, err := sql.Open("sqlite3", common.GetDbPath())
	if err != nil {
		log.Fatalf("Error open sqlite3: %v", err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Error load driver sqlite3: %v", err)
	}

	fSrc, err := (&file.File{}).Open("./quiz/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance(
		"file",
		fSrc,
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("start migration!\n")
	if isDrop {
		err = m.Drop()
		log.Printf("run drop!\n")
	} else {
		err = m.Up()
		log.Printf("run up!\n")
	}
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Error migrating Up: %v", err)
	}
	log.Printf("migratation done!\n")
}
