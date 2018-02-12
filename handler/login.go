package handler

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/session"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

type loginHandler struct {
	sessionStorage  session.SessionService
	customerStorage customer.CustomerService
}

// New creates a new login handler
func NewLoginHandler(session session.SessionService, customer customer.CustomerService) *loginHandler {
	return &loginHandler{sessionStorage: session, customerStorage: customer}
}

func (h *loginHandler) View(w http.ResponseWriter, r *http.Request) error {
	t := template.New("login.html")
	t.Parse(string(templates.MustAsset("login.html")))
	return t.Execute(w, nil)
}

func (h *loginHandler) Auth(w http.ResponseWriter, r *http.Request) (err error) {
	form := authForm{
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	cust, err := h.customerStorage.Get(form.Email)
	if err != nil {
		form.GeneralError = err.Error()
		t := template.New("login.html")
		t.Parse(string(templates.MustAsset("login.html")))
		return t.Execute(w, form)
	}

	if cust.Password != form.Password {
		form.GeneralError = "senha errada"
		t := template.New("login.html")
		t.Parse(string(templates.MustAsset("login.html")))
		return t.Execute(w, form)
	}

	redirect := "/login?redirect=/x"
	if r.FormValue("redirect") != "" {
		redirect = r.FormValue("redirect")
	}

	h.sessionStorage.Set(form.Email)
	http.Redirect(w, r, redirect, http.StatusFound)
	return
}

type authForm struct {
	Email        string
	Password     string
	FormErrors   map[string][]string
	GeneralError string
}
