package login

import (
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

var Path string = "/login"

type SessionService interface {
	Get(r *http.Request) string
	Set(w http.ResponseWriter, session string) error
}

type handler struct {
	session SessionService
}

func New(session SessionService) *handler {
	return &handler{session: session}
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

	p := "/perfil"
	if r.FormValue("redirect") != "" {
		p = r.FormValue("redirect")
	}

	session := h.session.Get(r)
	if session != "" {
		http.Redirect(w, r, p, http.StatusFound)
		return nil
	}

	t := template.New("login.html")
	t.Parse(string(templates.MustAsset("login.html")))
	return t.Execute(w, nil)
}

func (h *handler) auth(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()

	p := Path + "?redirect=/perfil"
	if r.FormValue("redirect") != "" {
		p = r.FormValue("redirect")
	}

	session := r.PostFormValue("email")

	h.session.Set(w, session)
	http.Redirect(w, r, p, http.StatusFound)
	return
}
