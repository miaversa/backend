package register

import (
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/cadastro"

var templateFile = "register.html"
var defaultRedirectPath = "/perfil"

type handler struct {
}

// New creates a new register handler
func New() *handler {
	return &handler{}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.view(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.register(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) view(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	return t.Execute(w, nil)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) (err error) {
	return
}

func validate(r *http.Request) (bool, map[string][]string) {
	return false, map[string][]string{}
}
