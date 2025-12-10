package requests

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var createTpl *template.Template

func InitCreateTemplate(t *template.Template) {
	createTpl = t
}

func CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	gameIDStr := r.URL.Query().Get("game_id")
	if gameIDStr == "" {
		http.Error(w, "game_id не указан", http.StatusBadRequest)
		return
	}

	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		http.Error(w, "некорректный game_id", http.StatusBadRequest)
		return
	}

	// Получаем название игры
	title, err := GetGameTitleByID(gameID)
	if err != nil {
		http.Error(w, "Игра не найдена", http.StatusNotFound)
		return
	}

	// Данные которые попадут в шаблон
	data := struct {
		GameID    int
		GameTitle string
	}{
		GameID:    gameID,
		GameTitle: title,
	}

	// Рендерим шаблон
	var buf bytes.Buffer
	err = createTpl.ExecuteTemplate(&buf, "create_request.html", data)
	if err != nil {
		log.Println("Ошибка рендеринга create_request:", err)
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}
