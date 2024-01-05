package main

import (
	// "database/sql"
	"errors"
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
	"quiz/internal/common"
)

func main() {
	fmt.Printf("run setup\n\n")
	runMigrate()
}

func runMigrate() {
	cfg := mysqlDriver.Config{
		User:                 common.GetDotEnvVariable("DBUSER"),
		Passwd:               common.GetDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 common.GetDotEnvVariable("DBADDR"),
		DBName:               common.GetDotEnvVariable("DBNAME"),
		AllowNativePasswords: true,
	}
	log.Printf("DSN %+v\n", cfg.FormatDSN())
	log.Printf("start migration!\n")
	m, err := migrate.New("file://quiz/migrations", fmt.Sprintf("mysql://%s", cfg.FormatDSN()))
	if err != nil {
		log.Fatalf("Error loading migrations: %v", err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Error migrating Up: %v", err)
	}
	log.Printf("migratation done!\n")
}
