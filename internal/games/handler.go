package games

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func InitTemplates(t *template.Template) {
	tpl = t
}

// GamesHandler отображает список игр
func GamesHandler(w http.ResponseWriter, r *http.Request) {
	games, err := GetGamesList()
	if err != nil {
		http.Error(w, "Ошибка получения игр", http.StatusInternalServerError)
		log.Println("Ошибка в GamesHandler:", err)
		return
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "games.html", games)
	if err != nil {
		log.Println("Ошибка шаблона:", err)
		http.Error(w, "Ошибка рендеринга шаблона", http.StatusInternalServerError)
		return
	}

	_, _ = buf.WriteTo(w)
}
