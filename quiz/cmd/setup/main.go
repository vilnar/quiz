package main

import (
	// "database/sql"
	"fmt"
	"errors"
	"github.com/joho/godotenv"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"os"
	"log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
)

func getDotEnvVariable(key string) string {
	err := godotenv.Load("quiz/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	fmt.Printf("run setup\n\n")
	runMigrate()
}


func runMigrate() {
	cfg := mysqlDriver.Config{
		User:                 getDotEnvVariable("DBUSER"),
		Passwd:               getDotEnvVariable("DBPASS"),
		Net:                  "tcp",
		Addr:                 getDotEnvVariable("DBADDR"),
		DBName:               getDotEnvVariable("DBNAME"),
		AllowNativePasswords: true,
	}
	fmt.Printf("debug dsn %+v\n", cfg.FormatDSN())
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
