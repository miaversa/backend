package handler

import (
	"github.com/miaversa/backend/customer"
	"github.com/miaversa/backend/session"
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

type paymentHandler struct {
	sessionService  session.SessionService
	customerService customer.CustomerService
}

// New creates a new payment handler
func NewPaymentHandler(session session.SessionService, customer customer.CustomerService) *paymentHandler {
	return &paymentHandler{session, customer}
}

func (h *paymentHandler) View(w http.ResponseWriter, r *http.Request) error {
	session, err := h.sessionService.Get()
	if err != nil {
		return err
	}

	if session == "" {
		http.Redirect(w, r, "/login?redirect=/pagar", http.StatusFound)
		return nil
	}

	_, err = h.customerService.GetShippingAddress(session)
	if err != nil {
		http.Redirect(w, r, "/envio?redirect=/pagar", http.StatusFound)
		return nil
	}

	t := template.New("payment.html")
	t.Parse(string(templates.MustAsset("payment.html")))
	return t.Execute(w, nil)
}

func (h *paymentHandler) pay(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}
