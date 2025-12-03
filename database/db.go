package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"prototypeZ/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect подключает БД с использованием конфигурации из config
func Connect(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	handleError(err, "Ошибка открытия БД")

	// Проверка соединения
	err = DB.Ping()
	handleError(err, "Ошибка подключения к БД")

	// Настройка пула соединений
	DB.SetMaxOpenConns(20)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxLifetime(1 * time.Hour)

	log.Println("Подключение к базе данных успешно")
}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}
