package games

import (
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

	err = tpl.ExecuteTemplate(w, "layout.html", games)
	if err != nil {
		http.Error(w, "Ошибка рендеринга шаблона", http.StatusInternalServerError)
		log.Println("Ошибка шаблона:", err)
	}
}
