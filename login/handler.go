package login

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/session"
	"github.com/miaversa/backend/templates"
	"github.com/thedevsaddam/govalidator"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/login"

// Default redirect
var DefaultRedirectPath = "/perfil"

var templateFile = "login.html"

type handler struct {
	session         session.SessionService
	customerService customer.CustomerService
}

// New creates a new login handler
func New(session session.SessionService, customer customer.CustomerService) *handler {
	return &handler{session: session, customerService: customer}
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
	redirect := DefaultRedirectPath
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
	form := authForm{
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

	cust, err := h.customerService.Get(form.Email)
	if err != nil {
		form.GeneralError = err.Error()
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, form)
	}

	if cust.Password != form.Password {
		form.GeneralError = "senha errada"
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, form)
	}

	redirect := Path + "?redirect=" + DefaultRedirectPath
	if r.FormValue("redirect") != "" {
		redirect = r.FormValue("redirect")
	}

	h.session.Set(w, form.Email)
	http.Redirect(w, r, redirect, http.StatusFound)
	return
}

func validate(r *http.Request) (bool, map[string][]string) {
	validationRules := govalidator.MapData{
		"email":    []string{"required", "email"},
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

type authForm struct {
	Email        string
	Password     string
	FormErrors   map[string][]string
	GeneralError string
}
