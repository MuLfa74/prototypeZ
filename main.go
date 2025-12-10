package main

import (
	"html/template"
	"log"
	"net/http"
	"prototypeZ/config"
	"prototypeZ/database"
	"prototypeZ/internal/auth"
	"prototypeZ/internal/games"
	"prototypeZ/internal/requests"
)

func main() {
	// Загружаем конфиг
	cfg := config.Load()

	// Подключаем БД
	database.Connect(cfg)
	db := database.DB
	if db == nil {
		log.Fatal("DB not initialized")
	}

	// Парсим шаблоны
	tpl := template.Must(template.ParseGlob("internal/templates/*.html"))

	// Инициализация модулей
	games.InitTemplates(tpl)
	requests.InitTemplates(tpl)

	// --- AUTH ---
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService, tpl)

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

	// Главная страница редирект на /games
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/games", http.StatusSeeOther)
	})

	// Роуты модулей
	http.HandleFunc("/games", games.GamesHandler)
	http.HandleFunc("/requests", requests.RequestsHandler)

	// Статика
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
