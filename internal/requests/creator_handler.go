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
	if r.Method == http.MethodGet {
		renderCreateForm(w, r)
		return
	}

	if r.Method == http.MethodPost {
		handleCreateSubmit(w, r)
		return
	}

	http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
}

func renderCreateForm(w http.ResponseWriter, r *http.Request) {
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

	title, err := GetGameTitleByID(gameID)
	if err != nil {
		http.Error(w, "Игра не найдена", http.StatusNotFound)
		return
	}

	data := struct {
		GameID    int
		GameTitle string
	}{
		GameID:    gameID,
		GameTitle: title,
	}

	var buf bytes.Buffer
	err = createTpl.ExecuteTemplate(&buf, "create_request.html", data)
	if err != nil {
		http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func handleCreateSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка формы", http.StatusBadRequest)
		return
	}

	gameID, _ := strconv.Atoi(r.URL.Query().Get("game_id"))
	userID := 1 // TODO: получить из сессии

	// Забираем данные из формы
	typeVal := r.FormValue("type") == "1"
	sexVal := r.FormValue("sex") == "1"

	age, _ := strconv.Atoi(r.FormValue("age"))

	req := Request{
		GameID:    gameID,
		UserID:    userID,
		Type:      typeVal,
		Purpose:   r.FormValue("purpose"),
		Sex:       sexVal,
		Age:       uint8(age),
		Contact:   r.FormValue("contact"),
		PrimeTime: r.FormValue("prime-time"),
	}

	err = CreateNewRequest(req)
	if err != nil {
		log.Println("Ошибка записи в БД:", err)
		http.Error(w, "Ошибка БД", http.StatusInternalServerError)
		return
	}

	// после создания → переходим на список заявок
	http.Redirect(w, r, "/requests?game_id="+strconv.Itoa(gameID), http.StatusSeeOther)
}
