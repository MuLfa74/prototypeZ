package users

import (
    "html/template"
    "net/http"
    "strconv"
    "strings"

    "github.com/go-chi/chi/v5"
)

type Handler struct {
    service   Service
    templates *template.Template
}

func NewHandler(service Service, tmpl *template.Template) *Handler {
    return &Handler{service: service, templates: tmpl}
}

func (h *Handler) RegisterRoutes(r chi.Router) {
    r.Get("/profile/{id}", h.GetProfile)
    r.Post("/profile/{id}", h.UpdateProfile)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    user, err := h.service.GetProfile(r.Context(), id)
    if err != nil {
        http.Error(w, "user not found", http.StatusNotFound)
        return
    }

    h.templates.ExecuteTemplate(w, "profile.html", user)
}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "invalid id", http.StatusBadRequest)
        return
    }

    r.ParseForm()

    ageInt, _ := strconv.Atoi(r.FormValue("age"))
    gender := r.FormValue("gender")
    sex := gender == "мужчина" || gender == "1"

    gamesList := strings.Split(r.FormValue("games"), ",")

    user := &User{
        ID:        id,
        Sex:       sex,
        Age:       uint8(ageInt),
        Contact:   r.FormValue("contact"),
        PrimeTime: r.FormValue("prime-time"),
        Games:     gamesList,
    }

    err = h.service.UpdateProfile(r.Context(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    http.Redirect(w, r, "/profile/"+idStr, http.StatusSeeOther)
}
