package main

import (
	"log"
	"os"

	"github.com/aniqaqill/go-ecom/config"
	"github.com/aniqaqill/go-ecom/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Initialize MySQL DB connection
	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Net:                  "tcp",
		Addr:                 config.Envs.DBAddres,
		DBName:               config.Envs.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Setup migration driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the migration
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Path to your migrations
		"mysql",                         // DB type
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Check the command argument (up/down)
	cmd := os.Args[len(os.Args)-1]
	switch cmd {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}

	log.Println("Migration completed.")
}
