package handler

/*
import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/session"
	"github.com/miaversa/backend/shipping"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/pagar"

var templateFile string = "payment.html"

type handler struct {
	sessionService  session.SessionService
	customerService customer.CustomerService
}

// New creates a new payment handler
func New(session session.SessionService, customer customer.CustomerService) *handler {
	return &handler{session, customer}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := h.view(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		if err := h.pay(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) view(w http.ResponseWriter, r *http.Request) error {
	session := h.sessionService.Get(r)
	if session == "" {
		http.Redirect(w, r, "/login?redirect="+Path, http.StatusFound)
		return nil
	}

	_, ok := h.customerService.GetShippingAddress(session)
	if !ok {
		http.Redirect(w, r, shipping.Path, http.StatusFound)
		return nil
	}

	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	return t.Execute(w, nil)
}

func (h *handler) pay(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}
*/
