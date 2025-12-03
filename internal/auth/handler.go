package auth

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type Handler struct {
	service   *Service
	templates *template.Template
}

func NewHandler(service *Service) *Handler {
	tmpl := template.Must(template.ParseGlob(filepath.Join("internal", "templates", "*.html")))
	return &Handler{service: service, templates: tmpl}
}

func (h *Handler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "login.html", nil)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "form error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := h.service.Login(email, password)
	if err != nil {
		h.templates.ExecuteTemplate(w, "login.html", map[string]any{
			"Error": "Неверный логин или пароль",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "user_id",
		Value: fmt.Sprintf("%d", user.ID),
		Path:  "/",
	})

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func (h *Handler) ShowRegister(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "register.html", nil)
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "form error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if err := h.service.Register(email, password, confirm); err != nil {
		h.templates.ExecuteTemplate(w, "register.html", map[string]any{
			"Error": err.Error(),
		})
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
