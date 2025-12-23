package games

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func InitTemplates(t *template.Template) {
	tpl = t
}

type PageData struct {
	UserID string
	Games  any
}

func GamesHandler(w http.ResponseWriter, r *http.Request) {
	games, err := GetGamesList()
	if err != nil {
		http.Error(w, "Ошибка получения игр", http.StatusInternalServerError)
		return
	}

	var userID string
	if c, err := r.Cookie("user_id"); err == nil {
		userID = c.Value
	}

	data := PageData{
		UserID: userID,
		Games:  games,
	}

	tpl.ExecuteTemplate(w, "games.html", data)
}
