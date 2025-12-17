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
	cfg := config.Load()

	database.Connect(cfg)
	db := database.DB
	if db == nil {
		log.Fatal("DB not initialized")
	}

	tpl := template.Must(template.ParseGlob("internal/templates/*.html"))

	games.InitTemplates(tpl)
	requests.InitTemplates(tpl)
	requests.InitCreateTemplate(tpl)

	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService, tpl)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowRegister(w, r)
			return
		}
		if r.Method == http.MethodPost {
			authHandler.HandleRegister(w, r)
			return
		}
		http.NotFound(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			authHandler.ShowLogin(w, r)
			return
		}
		if r.Method == http.MethodPost {
			authHandler.HandleLogin(w, r)
			return
		}
		http.NotFound(w, r)
	})

	// ПРОФИЛЬ
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		data := map[string]any{
			"UserID": c.Value,
		}

		err = tpl.ExecuteTemplate(w, "profile.html", data)
		if err != nil {
			http.Error(w, "Ошибка загрузки профиля", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/games", http.StatusSeeOther)
	})

	http.HandleFunc("/games", games.GamesHandler)
	http.HandleFunc("/requests", requests.RequestsHandler)
	http.HandleFunc("/requests/create_request", requests.CreateRequestHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
