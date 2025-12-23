package database

import (
	"database/sql"
	"fmt"
	"log"
	"prototypeZ/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
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
