package main

import (
	"html/template"
	"log"
	"net/http"
	"prototypeZ/config"
	"prototypeZ/database"
	"prototypeZ/internal/games"
	"prototypeZ/internal/requests"
)

func main() {
	// Загружаем конфигурацию
	cfg := config.Load()

	// Подключаем БД
	database.Connect(cfg)

	// Проверка работы
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	// Парсим шаблоны
	tpl := template.Must(template.ParseGlob("internal/templates/*.html"))
	games.InitTemplates(tpl)
	requests.InitTemplates(tpl)

	// Роут
	http.HandleFunc("/games", games.GamesHandler)
	http.HandleFunc("/requests", requests.RequestsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
