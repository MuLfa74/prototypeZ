package requests

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// tpl передаём из main.go
var tpl *template.Template

func InitTemplates(t *template.Template) {
	tpl = t
}

type RequestsPageData struct {
	GameTitle string
	Requests  []Request
}

func handleError(w http.ResponseWriter, statusCode int, message string, err error) {
	http.Error(w, message, statusCode)
	if err != nil {
		log.Println(err)
	}
}

func RequestsHandler(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.URL.Query().Get("game_id")
	if gameIDStr == "" {
		handleError(w, http.StatusBadRequest, "game_id не указан", nil)
		return
	}

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "Некорректный game_id", err)
		return
	}

	reqs, err := GetRequestsForGame(gameID)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Ошибка получения запросов", err)
		return
	}

	gameTitle, err := GetGameTitleByID(gameID)
	if err != nil {
		http.Error(w, "Игра не найдена", http.StatusNotFound)
		return
	}

	pageData := RequestsPageData{
		GameTitle: gameTitle,
		Requests:  reqs,
	}

	var buf bytes.Buffer
	err = tpl.ExecuteTemplate(&buf, "requests.html", pageData)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Ошибка рендеринга шаблона", err)
		return
	}

	_, _ = buf.WriteTo(w)
}
