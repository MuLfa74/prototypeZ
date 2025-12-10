package auth

import (
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct {
	service   *Service
	templates *template.Template
}

func NewHandler(s *Service, tpl *template.Template) *Handler {
	return &Handler{
		service:   s,
		templates: tpl,
	}
}

func (h *Handler) ShowLogin(w http.ResponseWriter, r *http.Request) {
	h.templates.ExecuteTemplate(w, "login.html", nil)
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "form error", http.StatusBadRequest)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	user, err := h.service.Login(login, password)
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

	login := r.FormValue("login")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if err := h.service.Register(login, password, confirm); err != nil {
		h.templates.ExecuteTemplate(w, "register.html", map[string]any{
			"Error": err.Error(),
		})
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
