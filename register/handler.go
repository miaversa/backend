package register

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/session"
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
	sessionService  session.SessionService
	customerService customer.CustomerService
}

// New creates a new register handler
func New(s session.SessionService, c customer.CustomerService) *handler {
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
	r.ParseForm()
	redirect := DefaultRedirectPath
	if r.FormValue("redirect") != "" {
		redirect = r.FormValue("redirect")
	}

	session := h.sessionService.Get(r)
	if session != "" {
		http.Redirect(w, r, redirect, http.StatusFound)
		return nil
	}

	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	return t.Execute(w, nil)
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	form := registerForm{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}

	valid, errors := validate(r)
	form.FormErrors = errors
	if !valid {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, form)
	}

	c := customer.Customer{
		Email:    form.Email,
		Name:     form.Name,
		Password: form.Password,
	}

	if err := h.customerService.Put(c); err != nil {
		form.GeneralError = err.Error()
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, form)
	}
	h.sessionService.Set(w, form.Email)
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

type registerForm struct {
	Name         string
	Email        string
	Password     string
	FormErrors   map[string][]string
	GeneralError string
}
