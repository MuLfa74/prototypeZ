package database

import (
	"database/sql"
	"log"
	"prototypeZ/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(cfg *config.Config) *sql.DB {
	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ")/" + cfg.DBName
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping DB:", err)
	}

	DB = db
	log.Println("Connected to DB")
	return db
}
