package main

import (
	"html/template"
	"log"
	"net/http"
	"prototypeZ/config"
	"prototypeZ/database"
	"prototypeZ/internal/games"
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
	// Games
	tpl := template.Must(template.ParseGlob("internal/templates/*.html"))
	games.InitTemplates(tpl)

	// Роут
	http.HandleFunc("/games", games.GamesHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
