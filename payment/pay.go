package payment

import (
	"github.com/miaversa/backend/templates"
	"html/template"
	"net/http"
)

// Path for the routing
var Path string = "/pagar"

var templateFile string = "payment.html"

type handler struct {
}

// New creates a new payment handler
func New() *handler {
	return &handler{}
}

func (s *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := template.New(templateFile)
	t.Parse(string(templates.MustAsset(templateFile)))
	t.Execute(w, nil)
}
