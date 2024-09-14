package main

import (
	"database/sql"
	"log"

	"github.com/aniqaqill/go-ecom/cmd/api"
	"github.com/aniqaqill/go-ecom/config"
	"github.com/aniqaqill/go-ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
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

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database is connected")
}
