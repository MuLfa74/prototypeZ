package main

import (
	"html/template"
	"log"
	"net/http"
	"prototypeZ/config"
	"prototypeZ/database"
	"prototypeZ/internal/auth"
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
	tpl := template.Must(template.ParseGlob("internal/templates/*.html"))

	// Games
	games.InitTemplates(tpl)

	// --- AUTH INIT ---
	authRepo := auth.NewRepository(database.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// --- AUTH ROUTES ---
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authHandler.ShowLogin(w, r)
		case http.MethodPost:
			authHandler.HandleLogin(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			authHandler.ShowRegister(w, r)
		case http.MethodPost:
			authHandler.HandleRegister(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	// Роут игр
	http.HandleFunc("/games", games.GamesHandler)

	// Static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
