package shipping

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/login"
	"github.com/miaversa/backend/templates"
	"github.com/thedevsaddam/govalidator"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/entrega"

var templateFile = "shipping.html"
var defaultRedirectPath = "/perfil"

type handler struct {
	sessionService  login.SessionService
	customerService customer.CustomerService
}

// New creates a new shipping handler
func New(s login.SessionService, c customer.CustomerService) *handler {
	return &handler{sessionService: s, customerService: c}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.view(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.shipping(w, r); err != nil {
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

func (h *handler) shipping(w http.ResponseWriter, r *http.Request) (err error) {

	session := h.sessionService.Get(r)
	if session == "" {
		http.Redirect(w, r, "/login?redirect=/shipping", http.StatusFound)
		return nil
	}

	r.ParseForm()
	street := r.PostFormValue("street")
	number := r.PostFormValue("number")
	complement := r.PostFormValue("complement")
	district := r.PostFormValue("district")
	city := r.PostFormValue("city")
	state := r.PostFormValue("state")
	country := r.PostFormValue("country")
	postalCode := r.PostFormValue("postalCode")

	valid, errors := validate(r)
	_ = errors
	if !valid {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}

	shipAddr := customer.ShippingAddress{
		Street:     street,
		Number:     number,
		Complement: complement,
		District:   district,
		City:       city,
		State:      state,
		Country:    country,
		PostalCode: postalCode,
	}

	if err := h.customerService.SetShippingAddress(session, shipAddr); err != nil {
		t := template.New(templateFile)
		t.Parse(string(templates.MustAsset(templateFile)))
		return t.Execute(w, nil)
	}
	http.Redirect(w, r, Path, http.StatusFound)

	return
}

func validate(r *http.Request) (bool, map[string][]string) {

	validationRules := govalidator.MapData{
		"street":     []string{"required"},
		"number":     []string{"required"},
		"complement": []string{"required"},
		"district":   []string{"required"},
		"city":       []string{"required"},
		"state":      []string{"required"},
		"country":    []string{"required"},
		"postalCode": []string{"required"},
	}

	validationMessages := govalidator.MapData{
		"street":     []string{"required: Endereço requerido"},
		"number":     []string{"required: Número requerido"},
		"complement": []string{"required: Complemento requerido"},
		"district":   []string{"required: Bairro requerido"},
		"city":       []string{"required: Cidade requerido"},
		"state":      []string{"required: Estado requerido"},
		"country":    []string{"required: País requerido"},
		"postalCode": []string{"required: CEP requerido"},
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
