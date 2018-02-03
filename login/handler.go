package login

import (
	"github.com/miaversa/backend/templates"
	"github.com/thedevsaddam/govalidator"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/login"

var templateFile = "login.html"
var defaultRedirectPath = "/perfil"

// SessionService defines the session api
type SessionService interface {
	Get(r *http.Request) string
	Set(w http.ResponseWriter, session string) error
}

type AuthService interface {
	Validate(email, password string) bool
}

type handler struct {
	session     SessionService
	authService AuthService
}

// New creates a new login handler
func New(session SessionService, auth AuthService) *handler {
	return &handler{session: session, authService: auth}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.view(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.auth(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) view(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	redirect := defaultRedirectPath
	if r.FormValue("redirect") != "" {
		redirect = r.FormValue("redirect")
	}

	session := h.session.Get(r)
	if session != "" {
		http.Redirect(w, r, redirect, http.StatusFound)
		return nil
	}

	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	return t.Execute(w, nil)
}

func (h *handler) auth(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()

	valid, errors := validate(r)
	_ = errors
	if !valid {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}

	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	if !h.authService.Validate(email, password) {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}

	redirect := Path + "?redirect=" + defaultRedirectPath
	if r.FormValue("redirect") != "" {
		redirect = r.FormValue("redirect")
	}

	h.session.Set(w, email)
	http.Redirect(w, r, redirect, http.StatusFound)
	return
}

func validate(r *http.Request) (bool, map[string][]string) {
	validationRules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"password": []string{"required", "between:6,20"},
	}

	validationMessages := govalidator.MapData{
		"email":    []string{"required:email requerido"},
		"password": []string{"required:a senha é requerida.", "between:a senha deve ter no minimo 6 caracteres"},
	}

	v := govalidator.New(govalidator.Options{
		Request:         r,
		Rules:           validationRules,
		Messages:        validationMessages,
		RequiredDefault: true,
	})

	errors := v.Validate()
	return len(errors) == 0, errors
}
