package register

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/login"
	"github.com/miaversa/backend/templates"
	"github.com/thedevsaddam/govalidator"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/cadastro"
var DefaultRedirectPath = "/pagar"

var templateFile = "register.html"

type handler struct {
	sessionService  login.SessionService
	customerService customer.CustomerService
}

// New creates a new register handler
func New(s login.SessionService, c customer.CustomerService) *handler {
	return &handler{sessionService: s, customerService: c}
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
	session := h.sessionService.Get(r)
	if session != "" {
		http.Redirect(w, r, DefaultRedirectPath, http.StatusFound)
		return nil
	}

	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	return t.Execute(w, nil)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	email := r.PostFormValue("email")
	name := r.PostFormValue("name")
	password := r.PostFormValue("password")

	valid, errors := validate(r)
	_ = errors
	if !valid {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}

	c := customer.Customer{
		Email:    email,
		Name:     name,
		Password: password,
	}

	if err := h.customerService.Put(c); err != nil {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}
	h.sessionService.Set(w, email)
	http.Redirect(w, r, Path, http.StatusFound)

	return
}

func validate(r *http.Request) (bool, map[string][]string) {
	validationRules := govalidator.MapData{
		"email":    []string{"required", "email"},
		"name":     []string{"required"},
		"password": []string{"required"},
	}

	validationMessages := govalidator.MapData{
		"email":    []string{"required:email requerido."},
		"password": []string{"required:a senha é requerida."},
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
